package controller

import (
	"context"
	"goods/dao/redis"
	"goods/utils"
	"google.golang.org/grpc"
	"log"
	"time"
)

func VisitInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("before visitInterceptor. req: %+v", req)
	// 最近1天的访问量, 设置1天的超时时间
	recentVisitKey := "goods_visit_" + info.FullMethod
	res := redis.RDB.Incr(recentVisitKey)
	utils.CheckErr(res.Err(), "redis incr failed, keyName:"+recentVisitKey)
	redis.RDB.Expire(recentVisitKey, time.Hour*24)

	// 历史访问量，不设超时时间
	historyVisitKey := "goods_visit_history_" + info.FullMethod
	res = redis.RDB.Incr(historyVisitKey)
	utils.CheckErr(res.Err(), "redis incr failed, keyName:"+historyVisitKey)
	log.Printf("after visitInterceptor")

	resp, err := handler(ctx, req)
	return resp, err
}
