package v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"order/conf"
	"order/models"
	"order/pb"
	"order/utils"
	"time"
)

const CacheKeyPrefix = "order_cache_"

type OrderServer struct{}

var goodsClient pb.GoodsClient

func InitGoodsRPCClient() {
	conn, err := grpc.Dial(conf.Config.GoodsServiceAddr, grpc.WithInsecure())
	utils.Fatalf(err, "")
	goodsClient = pb.NewGoodsClient(conn)
}

func (s *OrderServer) CreateOrder(ctx context.Context, request *pb.OrderRequest) (*pb.OrderCommonReply, error) {
	if len(request.GoodsUuid)==0{
		msg := "goodsUuid 不允许为空"
		return &pb.OrderCommonReply{Msg:msg}, errors.New(msg)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	goodsRequest := pb.GoodsRequest{GoodsUuid:request.GoodsUuid}
	replyStockDown, err := goodsClient.MakeStockDown(ctx, &goodsRequest)
	err = utils.CheckRPCError(err)
	if err != nil {
		return &pb.OrderCommonReply{Msg:replyStockDown.Msg}, err
	}
	item := models.Order{
		GoodsUuid:request.GoodsUuid,
		GoodsTypeId:request.GoodsTypeId,
		PrimaryType:request.PrimaryType,
		SecondaryType:request.SecondaryType,
		Price:request.Price,
		Title:request.Title,
		Subtitle:request.Subtitle,
		OrderStatus:"new",
	}
	reply, err := models.CreateOrder(&item)
	if err != nil {
		utils.Logf(err, "")
		// TODO: 临时方案，应当使用分布式事务
		replyStockUp, err := goodsClient.MakeStockUp(ctx, &goodsRequest)
		err = utils.CheckRPCError(err)
		if err != nil {
			return &pb.OrderCommonReply{Msg:replyStockUp.Msg}, err
		}
		return reply, status.Errorf(codes.InvalidArgument, "检查你的参数:%+v", request)
	}
	return reply, err
}

func (s *OrderServer) GetOrderList(ctx context.Context, request *pb.OrderRequest) (*pb.OrderReply, error) {
	// 如果缓存存在，就取出来返回结果
	cache := utils.Cache{
		RedisKeyName: fmt.Sprintf(
			"%sorderList_%s_%s_%s_%d_%d",
			CacheKeyPrefix,
			request.GoodsUuid,
			request.PrimaryType,
			request.SecondaryType,
			request.PageIndex,
			request.PageSize,
		),
	}
	if resp, err := cache.GetStringCache(); err == nil {
		var model pb.OrderReply
		if err := json.Unmarshal([]byte(resp), &model); err != nil {
			utils.Logf(err, "redis json unmarshal failed, keyName:"+cache.RedisKeyName)
			return nil, errors.New("redis data unmarshal error")
		}
		return &model, nil
	} else {
		reply, err := models.QueryOrder(request)
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

func (s *OrderServer) GetOrderStatistic(ctx context.Context, request *pb.OrderRequest) (*pb.OrderStatisticReply, error) {
	// 如果缓存存在，就取出来返回结果
	cache := utils.Cache{
		RedisKeyName: fmt.Sprintf(
			"%sorderStatistic",
			CacheKeyPrefix,
		),
	}
	if resp, err := cache.GetStringCache(); err == nil {
		var model pb.OrderStatisticReply
		if err := json.Unmarshal([]byte(resp), &model); err != nil {
			utils.Logf(err, "redis json unmarshal failed, keyName:"+cache.RedisKeyName)
			return nil, errors.New("redis data unmarshal error")
		}
		return &model, nil
	} else {
		reply := models.GetOrderStatistic()
		cache.Result = *reply
		// 设置缓存
		if err := cache.StoreStringCache(); err != nil {
			utils.Logf(err, "redis set cache failed")
			return reply, err
		}
		return reply, nil
	}
}
