package main

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"log"
	"net"
	"order/conf"
	"order/controller/v1"
	"order/models"
	"order/pb"
	"order/utils"
)

func runServer() {
	conf.InitConfig()
	models.InitGorm()
	utils.InitRedis()
	v1.InitGoodsRPCClient()

	lis, err := net.Listen("tcp", ":"+conf.Config.AppPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				utils.RequestParamInterceptor,
			)))
	pb.RegisterOrderServer(s, &v1.OrderServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	runServer()
}
