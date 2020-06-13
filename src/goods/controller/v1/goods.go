package v1

import (
	"context"
	"errors"
	"fmt"
	"goods/models"
	"goods/pb"
	"goods/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GoodsServer struct{}

func (s *GoodsServer) GetGoodsList(ctx context.Context, request *pb.GoodsRequest) (*pb.GoodsReply, error) {
	// 如果缓存存在，就取出来返回结果
	// 如果没有缓存再需要后续操作，取得service层的结果，返回响应并set到context，会在缓存中间件把响应缓存起来
	if value, ok := utils.InterceptorKeyGoodsList[utils.CacheKeyPrefix+"/pb.Goods/GetGoodsList"]; ok {
		// 命中缓存时,直接return
		return value, nil
	} else {
		reply, err := models.GetGoodsByType(request)
		if err != nil {
			utils.Logf(err, "")
			return nil, status.Errorf(codes.InvalidArgument, "检查你的参数:%+v", request)
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
	if value, ok := utils.InterceptorKeyGoodsDetail[utils.CacheKeyPrefix+"/pb.Goods/GetGoodsDetail"]; ok {
		// 命中缓存时, 更新商品热度，再return
		if err := updateHotIndex(value); err != nil {
			utils.Logf(err, "")
			return value, err
		}
		return value, nil
	} else {
		reply, err := models.GetGoods(request)
		if err != nil {
			utils.Logf(err, "")
			return nil, status.Errorf(codes.InvalidArgument, "检查你的参数:%+v", request)
		} else {
			// 更新商品热度
			if err := updateHotIndex(reply); err != nil {
				utils.Logf(err, "")
				return value, err
			}
		}
		return reply, nil
	}
}
