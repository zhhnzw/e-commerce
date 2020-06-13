package utils

import (
	"context"
	"goods/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func RequestParamInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
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
