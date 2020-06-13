package v1

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"goods/conf"
	"goods/models"
	"goods/pb"
	"goods/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type goodsServer struct{}

func (s *goodsServer) GetGoodsList(ctx context.Context, request *pb.GoodsRequest) (*pb.GoodsReply, error) {
	reply, err := models.GetGoodsByType(request)
	if err != nil {
		utils.Logf(err, "")
		return nil, status.Errorf(codes.InvalidArgument, "检查你的参数:%+v", request)
	}
	return reply, nil
}

func (s *goodsServer) GetGoodsDetail(ctx context.Context, request *pb.GoodsRequest) (*pb.GoodsReplyItem, error) {
	reply, err := models.GetGoods(request)
	if err != nil {
		utils.Logf(err, "")
		return nil, status.Errorf(codes.InvalidArgument, "检查你的参数:%+v", request)
	}
	return reply, nil
}

func Run() {
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
	pb.RegisterGoodsServer(s, &goodsServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
