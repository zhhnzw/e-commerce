package v1

import (
	"context"
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
	reply, err := models.GetCommoditiesByType(request)
	if err != nil {
		utils.Logf(err, "")
		return nil, status.Errorf(codes.InvalidArgument, "检查你的参数:%+v", request)
	}
	return reply, nil
}

func (s *goodsServer) GetGoodsDetail(ctx context.Context, request *pb.GoodsRequest) (*pb.GoodsReplyItem, error) {
	reply, err := models.GetCommodity(request)
	if err != nil {
		utils.Logf(err, "")
		return nil, status.Errorf(codes.InvalidArgument, "检查你的参数:%+v", request)
	}
	return reply, nil
}

func UnaryServerRequestParamInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("before handling. req: %+v", req)
	if request, ok := req.(*pb.GoodsRequest); ok {
		if request.PageIndex < 1 || request.PageSize < 1 || request.PageSize > 50 {
			return "", status.Errorf(codes.InvalidArgument, "pageIndex或pageSize参数错误,参数:%+v", request)
		}
	}
	resp, err := handler(ctx, req)
	log.Printf("after handling. resp: %+v", resp)
	return resp, err
}

func Run() {
	lis, err := net.Listen("tcp", ":"+conf.Config.AppPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(UnaryServerRequestParamInterceptor))
	pb.RegisterGoodsServer(s, &goodsServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
