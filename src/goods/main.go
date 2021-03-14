package main

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"goods/controller"
	"goods/controller/v1"
	"goods/dao/mysql"
	"goods/dao/redis"
	"goods/logger"
	"goods/pb"
	"goods/settings"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	if err := settings.Init(); err != nil {
		log.Printf("init settings failed, err:%v\n", err)
		return
	}
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		log.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	if err := mysql.InitGorm(settings.Conf.MySQLConfig); err != nil {
		log.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		log.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				controller.RequestParamInterceptor,
				controller.VisitInterceptor)))
	pb.RegisterGoodsServer(s, &v1.GoodsServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
