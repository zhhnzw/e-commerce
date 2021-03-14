package v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	goRedis "github.com/go-redis/redis"
	"go.uber.org/zap"
	"goods/dao/mysql"
	"goods/dao/redis"
	"goods/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

const CacheKeyPrefix = "goods_cache_"

type GoodsServer struct{}

func (s *GoodsServer) GetGoodsList(ctx context.Context, request *pb.GoodsRequest) (*pb.GoodsReply, error) {
	// 如果缓存存在，就取出来返回结果
	cache := redis.Cache{
		RedisKeyName: fmt.Sprintf(
			"%sgoodsList_%s_%s_%s_%d_%d",
			CacheKeyPrefix,
			request.GoodsUuid,
			request.PrimaryType,
			request.SecondaryType,
			request.PageIndex,
			request.PageSize,
		),
	}
	if resp, err := cache.GetStringCache(); err == nil {
		var model pb.GoodsReply
		if err := json.Unmarshal([]byte(resp), &model); err != nil {
			zap.L().Error("redis json unmarshal failed", zap.Error(err), zap.String("redisKey", cache.RedisKeyName))
			return nil, errors.New("redis data unmarshal error")
		}
		return &model, nil
	} else {
		reply, err := mysql.QueryGoods(request)
		if err != nil {
			zap.L().Error("", zap.Error(err))
			return reply, status.Errorf(codes.InvalidArgument, "你的参数:%+v \n returned err:%s", request, err.Error())
		}
		cache.Result = *reply
		// 设置缓存
		if err := cache.StoreStringCache(); err != nil {
			zap.L().Error("redis set cache failed", zap.Error(err))
			return reply, err
		}
		return reply, nil
	}
}

// 更新热门商品的热度: 修改商品类别下的访问访问量
func updateHotIndex(model *pb.GoodsReplyItem) error {
	if len(model.PrimaryType) == 0 || len(model.SecondaryType) == 0 {
		return errors.New(fmt.Sprintf("model:%+v", model))
	}
	hotGoodsPrimaryTypeKey := fmt.Sprintf("goods_hot_primary_type_%s_zset", model.PrimaryType)
	hotGoodsSecondaryTypeKey := fmt.Sprintf("goods_hot_secondary_type_%s_zset", model.SecondaryType)
	redis.RDB.ZIncrBy(hotGoodsPrimaryTypeKey, 1, model.GoodsUuid)
	redis.RDB.ZIncrBy(hotGoodsSecondaryTypeKey, 1, model.GoodsUuid)
	return nil
}

func (s *GoodsServer) GetGoodsDetail(ctx context.Context, request *pb.GoodsRequest) (*pb.GoodsReplyItem, error) {
	cache := redis.Cache{
		RedisKeyName: fmt.Sprintf("%sgoodsDetail_%s", CacheKeyPrefix, request.GoodsUuid),
	}
	if resp, err := cache.GetStringCache(); err == nil {
		// 命中缓存时, 更新商品热度，再return
		var goodsModel pb.GoodsReplyItem
		if err := json.Unmarshal([]byte(resp), &goodsModel); err != nil {
			zap.L().Error("redis json unmarshal failed", zap.Error(err), zap.String("redisKey", cache.RedisKeyName))
			return nil, errors.New("redis data unmarshal error")
		}
		if err := updateHotIndex(&goodsModel); err != nil {
			zap.L().Error("", zap.Error(err))
			return &goodsModel, err
		}
		return &goodsModel, nil
	} else {
		reply, err := mysql.GetGoods(request)
		if err != nil {
			zap.L().Error("", zap.Error(err))
			return reply, status.Errorf(codes.InvalidArgument, "你的参数:%+v \n returned err:%s", request, err.Error())
		} else {
			// 更新商品热度
			if err := updateHotIndex(reply); err != nil {
				zap.L().Error("", zap.Error(err))
				return reply, err
			}
		}
		cache.Result = *reply
		// 设置缓存
		if err := cache.StoreStringCache(); err != nil {
			zap.L().Error("redis set cache failed", zap.Error(err))
			return reply, err
		}
		return reply, nil
	}
}

func (s *GoodsServer) GetGoodsHotList(ctx context.Context, request *pb.GoodsRequest) (*pb.GoodsReply, error) {
	if len(request.PrimaryType) > 0 {
		hotCommoditiesPrimaryTypeKey := fmt.Sprintf("goods_hot_primary_type_%s_zset", request.PrimaryType)
		cmd := redis.RDB.ZRevRange(hotCommoditiesPrimaryTypeKey, request.PageIndex-1, request.PageSize)
		if res, err := cmd.Result(); err != nil {
			zap.L().Error("redis read key failed", zap.Error(err), zap.String("redisKey", hotCommoditiesPrimaryTypeKey))
		} else {
			// 取得商品榜单有序列表之后，再遍历取得详细数据
			results := make([]*pb.GoodsReplyItem, 0, len(res))
			for _, uuid := range res {
				// 优先取缓存
				cache := redis.Cache{
					RedisKeyName: fmt.Sprintf("%sgoodsDetail_%s", CacheKeyPrefix, uuid),
				}
				if r, err := cache.GetStringCache(); err == nil {
					log.Println("get data from cache:" + r)
					var goodsModel pb.GoodsReplyItem
					if err := json.Unmarshal([]byte(r), &goodsModel); err != nil {
						zap.L().Error("redis json unmarshal failed", zap.Error(err), zap.String("redisKey", cache.RedisKeyName))
					}
					results = append(results, &goodsModel)
				} else if err == goRedis.Nil {
					// 若缓存没有再读数据库
					query := pb.GoodsRequest{GoodsUuid: uuid}
					if resultModel, e := mysql.GetGoods(&query); e != nil {
						if err.Error() != "record not found" {
							zap.L().Error("数据库查询操作异常", zap.Error(err))
						} else {
							log.Println("tb_commodity record not found ---> uuid:" + uuid)
						}
					} else {
						results = append(results, resultModel)
					}
				} else {
					zap.L().Error("", zap.Error(err))
					return nil, errors.New("redis error")
				}
			}
			return &pb.GoodsReply{Data: results}, nil
		}
	}
	return nil, nil
}

func (s *GoodsServer) GetGoodsStatistic(ctx context.Context, request *pb.GoodsRequest) (*pb.GoodsStatisticReply, error) {
	// 如果缓存存在，就取出来返回结果
	cache := redis.Cache{
		RedisKeyName: fmt.Sprintf(
			"%sgoodsStatistic",
			CacheKeyPrefix,
		),
	}
	if resp, err := cache.GetStringCache(); err == nil {
		var model pb.GoodsStatisticReply
		if err := json.Unmarshal([]byte(resp), &model); err != nil {
			zap.L().Error("redis json unmarshal failed", zap.Error(err), zap.String("redisKey", cache.RedisKeyName))
			return nil, errors.New("redis data unmarshal error")
		}
		return &model, nil
	} else {
		reply := mysql.GetGoodsStatistic()
		cache.Result = *reply
		// 设置缓存
		if err := cache.StoreStringCache(); err != nil {
			zap.L().Error("redis set cache failed", zap.Error(err))
			return reply, err
		}
		return reply, nil
	}
}

func (s *GoodsServer) MakeStockUp(ctx context.Context, request *pb.GoodsRequest) (*pb.CommonReply, error) {
	return mysql.MakeStockUp(request)
}

func (s *GoodsServer) MakeStockDown(ctx context.Context, request *pb.GoodsRequest) (*pb.CommonReply, error) {
	return mysql.MakeStockDown(request)
}
