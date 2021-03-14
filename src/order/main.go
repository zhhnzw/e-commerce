package main

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"order/controller"
	"order/controller/v1"
	"order/dao/mysql"
	"order/dao/redis"
	"order/logger"
	"order/pb"
	"order/settings"
)

func runServer() {
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
				controller.RequestParamInterceptor)))
	pb.RegisterOrderServer(s, &v1.OrderServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	runServer()
}
