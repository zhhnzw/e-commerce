package v1_test

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"goods/pb"
	"goods/utils"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

const address = "localhost:50051"

func Test(t *testing.T) {
	Convey("Test", t, func() {
		// Set up a connection to the server.
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		goodsClient := pb.NewGoodsClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		Convey("Test GetGoodsStatistic", func() {
			request := &pb.GoodsRequest{}
			reply, err := goodsClient.GetGoodsStatistic(ctx, request)
			So(err, ShouldBeNil)
			utils.CheckRPCError(err)
			log.Printf("%+v\n", reply)
		})
		Convey("Test MakeStockUp", func() {
			request := &pb.GoodsRequest{GoodsUuid: "b7b10c01-62b8-42c7-a8c4-8efe119cd326"}
			reply, err := goodsClient.MakeStockUp(ctx, request)
			So(err, ShouldBeNil)
			utils.CheckRPCError(err)
			log.Printf("%+v\n", reply)
		})
	})
}
