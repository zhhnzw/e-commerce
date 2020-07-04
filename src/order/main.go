package main

import (
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"order/conf"
	"order/controller/v1"
	"order/models"
	"order/pb"
	"order/utils"
	"google.golang.org/grpc"
	"log"
	"net"
)

func runServer() {
	conf.InitConfig()
	utils.InitRedis()
	models.InitGorm()
	v1.InitGoodsRPCClient()
	gin.SetMode(gin.DebugMode)

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
