package main

import (
"context"
"goods/pb"
"google.golang.org/grpc"
"google.golang.org/grpc/codes"
"google.golang.org/grpc/status"
"log"
"reflect"
"time"
)

const address     = "localhost:50051"

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	goodsClient := pb.NewGoodsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	request := &pb.GoodsRequest{PageSize:10, PageIndex:1, PrimaryType:"clothes", SecondaryType:"shirt"}
	reply, err := goodsClient.GetGoodsList(ctx, request)
	if err != nil {
		s := status.Convert(err)
		switch s.Code() {
		case codes.InvalidArgument:
			log.Printf("InvalidArgument: %v", s)
		default:
			log.Printf("Unexpected type: %v", s)
		}
		log.Fatalf("could not greet: %v %s", err, reflect.TypeOf(err))
	}
	log.Printf("%+v\n", reply)
	request.GoodsUuid = "80e5d628-7762-4c80-89d3-3ea7c1486627"
	reply1, err := goodsClient.GetGoodsDetail(ctx, request)
	if err != nil {
		s := status.Convert(err)
		switch s.Code() {
		case codes.InvalidArgument:
			log.Printf("InvalidArgument: %v", s)
		default:
			log.Printf("Unexpected type: %v", s)
		}
		log.Fatalf("could not greet: %v %s", err, reflect.TypeOf(err))
	}
	log.Printf("%+v\n", reply1)
}


