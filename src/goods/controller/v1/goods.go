package v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"goods/models"
	"goods/pb"
	"goods/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

const CacheKeyPrefix = "goods_cache_"

type GoodsServer struct{}

func (s *GoodsServer) GetGoodsList(ctx context.Context, request *pb.GoodsRequest) (*pb.GoodsReply, error) {
	// 如果缓存存在，就取出来返回结果
	cache := utils.Cache{
		RedisKeyName: fmt.Sprintf("%sgoodsList_%s", CacheKeyPrefix, request.GoodsUuid),
	}
	if resp, err := cache.GetStringCache(); err == nil {
		var model pb.GoodsReply
		if err := json.Unmarshal([]byte(resp), &model); err != nil {
			utils.Logf(err, "redis json unmarshal failed, keyName:"+cache.RedisKeyName)
			return nil, errors.New("redis data unmarshal error")
		}
		return &model, nil
	} else {
		reply, err := models.GetGoodsByType(request)
		if err != nil {
			utils.Logf(err, "")
			return nil, status.Errorf(codes.InvalidArgument, "检查你的参数:%+v", request)
		}
		cache.Result = *reply
		// 设置缓存
		if err := cache.StoreStringCache(); err != nil {
			utils.Logf(err, "redis set cache failed")
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
	utils.RedisClient.ZIncrBy(hotGoodsPrimaryTypeKey, 1, model.GoodsUuid)
	utils.RedisClient.ZIncrBy(hotGoodsSecondaryTypeKey, 1, model.GoodsUuid)
	return nil
}

func (s *GoodsServer) GetGoodsDetail(ctx context.Context, request *pb.GoodsRequest) (*pb.GoodsReplyItem, error) {
	cache := utils.Cache{
		RedisKeyName: fmt.Sprintf("%sgoodsDetail_%s", CacheKeyPrefix, request.GoodsUuid),
	}
	if resp, err := cache.GetStringCache(); err == nil {
		// 命中缓存时, 更新商品热度，再return
		var goodsModel pb.GoodsReplyItem
		if err := json.Unmarshal([]byte(resp), &goodsModel); err != nil {
			utils.Logf(err, "redis json unmarshal failed, keyName:"+cache.RedisKeyName)
			return nil, errors.New("redis data unmarshal error")
		}
		if err := updateHotIndex(&goodsModel); err != nil {
			utils.Logf(err, "")
			return &goodsModel, err
		}
		return &goodsModel, nil
	} else {
		reply, err := models.GetGoods(request)
		if err != nil {
			utils.Logf(err, "")
			return nil, status.Errorf(codes.InvalidArgument, "检查你的参数:%+v", request)
		} else {
			// 更新商品热度
			if err := updateHotIndex(reply); err != nil {
				utils.Logf(err, "")
				return reply, err
			}
		}
		cache.Result = *reply
		// 设置缓存
		if err := cache.StoreStringCache(); err != nil {
			utils.Logf(err, "redis set cache failed")
			return reply, err
		}
		return reply, nil
	}
}

func (s *GoodsServer) GetGoodsHotList(ctx context.Context, request *pb.GoodsRequest) (*pb.GoodsReply, error) {
	if len(request.PrimaryType) > 0 {
		hotCommoditiesPrimaryTypeKey := fmt.Sprintf("goods_hot_primary_type_%s_zset", request.PrimaryType)
		cmd := utils.RedisClient.ZRevRange(hotCommoditiesPrimaryTypeKey, request.PageIndex-1, request.PageSize)
		if res, err := cmd.Result(); err != nil {
			utils.Logf(err, "redis read key failed, keyName:"+hotCommoditiesPrimaryTypeKey)
		} else {
			// 取得商品榜单有序列表之后，再遍历取得详细数据
			results := make([]*pb.GoodsReplyItem, 0, len(res))
			for _, uuid := range res {
				// 优先取缓存
				cache := utils.Cache{
					RedisKeyName: fmt.Sprintf("%sgoodsDetail_%s", CacheKeyPrefix, uuid),
				}
				if r, err := cache.GetStringCache(); err == nil {
					log.Println("get data from cache:" + r)
					var goodsModel pb.GoodsReplyItem
					if err := json.Unmarshal([]byte(r), &goodsModel); err != nil {
						utils.Logf(err, "redis json unmarshal failed, keyName:"+cache.RedisKeyName)
					}
					results = append(results, &goodsModel)
				} else if err == redis.Nil {
					// 若缓存没有再读数据库
					query := pb.GoodsRequest{GoodsUuid: uuid}
					if resultModel, e := models.GetGoods(&query); e != nil {
						if err.Error() != "record not found" {
							utils.Logf(err, "数据库查询操作异常")
						} else {
							log.Println("tb_commodity record not found ---> uuid:" + uuid)
						}
					} else {
						results = append(results, resultModel)
					}
				} else {
					utils.Logf(err, "")
					return nil, errors.New("redis error")
				}
			}
			return &pb.GoodsReply{Data: results}, nil
		}
	}
	return nil, nil
}
