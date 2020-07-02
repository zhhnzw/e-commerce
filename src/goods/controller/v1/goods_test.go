package v1_test

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"goods/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"reflect"
	"testing"
	"time"
)

const address = "localhost:50051"

func Test(t *testing.T) {
	Convey("TestGetCommodities", t, func() {
		// Set up a connection to the server.
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		goodsClient := pb.NewGoodsClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		request := &pb.GoodsRequest{}
		reply, err := goodsClient.GetGoodsStatistic(ctx, request)
		So(err, ShouldBeNil)
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
	})
}
