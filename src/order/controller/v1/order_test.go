package v1_test

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"google.golang.org/grpc"
	"log"
	"order/pb"
	"order/utils"
	"testing"
	"time"
)

const address = "localhost:50052"

func Test(t *testing.T) {
	Convey("Test", t, func() {
		// Set up a connection to the server.
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		orderClient := pb.NewOrderClient(conn)
		Convey("Test GetOrderStatistic", func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			request := &pb.OrderRequest{}
			reply, err := orderClient.GetOrderStatistic(ctx, request)
			So(err, ShouldBeNil)
			utils.CheckRPCError(err)
			log.Printf("%+v\n", reply)
		})
		Convey("Test GetOrderList", func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			request := &pb.OrderRequest{PageIndex: 1, PageSize: 10}
			reply, err := orderClient.GetOrderList(ctx, request)
			So(err, ShouldBeNil)
			utils.CheckRPCError(err)
			log.Printf("%+v\n", reply)
		})
		Convey("Test CreateOrder", func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			request := &pb.OrderRequest{
				GoodsUuid:     "b7b10c01-62b8-42c7-a8c4-8efe119cd326",
				PrimaryType:   "clothes",
				SecondaryType: "shirt",
				GoodsTypeId:   1,
				Price:         1931,
				Title:         "CYGNENOIR 20AW秋季字母印花宽松迷彩衬衫日系复古街头衬衣男潮",
				Subtitle:      "衬衣 subtitle",
				Img:           "http://shihuo.hupucdn.com/def/20200430/390bb16aeb93924ccd4c67f5f178cf9a1588213097.jpg",
				UserName:      "user",
				NickName:      "卖萌小公举",
				Mobile:        "18321212121",
				Email:         "2131@11.com",
			}
			reply, err := orderClient.CreateOrder(ctx, request)
			log.Printf("%+v\n", reply)
			So(err, ShouldBeNil)
			utils.CheckRPCError(err)
		})
	})
}
