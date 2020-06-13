package main

import (
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"goods/conf"
	"goods/controller/v1"
	"goods/models"
	"goods/pb"
	"goods/utils"
	"google.golang.org/grpc"
	"log"
	"net"
)

func runServer() {
	conf.InitConfig()
	utils.InitRedis()
	models.InitGorm()
	gin.SetMode(gin.DebugMode)

	lis, err := net.Listen("tcp", ":"+conf.Config.AppPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				utils.RequestParamInterceptor,
				utils.VisitInterceptor,
				utils.CacheInterceptor)))
	pb.RegisterGoodsServer(s, &v1.GoodsServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	runServer()
}
