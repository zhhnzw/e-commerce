// 随机新建订单
package main

import (
	"context"
	"fmt"
	"github.com/robfig/cron"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"order/conf"
	"order/models"
	"order/pb"
	"order/utils"
	"time"
)

var goodsClient pb.GoodsClient

func InitGoodsRPCClient() {
	conn, err := grpc.Dial(conf.Config.GoodsServiceAddr, grpc.WithInsecure())
	utils.Fatalf(err, "")
	goodsClient = pb.NewGoodsClient(conn)
}

var orderClient pb.OrderClient

func InitOrderRPCClient() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	utils.Fatalf(err, "")
	orderClient = pb.NewOrderClient(conn)
}

func run() {
	models.InitGorm()
	InitGoodsRPCClient()
	InitOrderRPCClient()
	userNameSrc := []string{"26394826", "127308924", "381936289", "434247532", "534545642"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	// 随机选取一个goodsUuid
	// 设置不同的seed，后续的随机操作才是真随机
	rand.Seed(time.Now().UnixNano())
	randId := rand.Int63n(10000000)
	goodsRequest := pb.GoodsRequest{PageSize: 1, PageIndex: randId}
	reply, err := goodsClient.GetGoodsList(ctx, &goodsRequest)
	log.Printf("randId:%d goods reply:%+v", randId, reply)
	if err == nil && len(reply.Data) > 0 {
		model := pb.OrderRequest{
			GoodsUuid:     reply.Data[0].GoodsUuid,
			PrimaryType:   reply.Data[0].PrimaryType,
			SecondaryType: reply.Data[0].SecondaryType,
			GoodsTypeId:   reply.Data[0].GoodsTypeId,
			Title:         reply.Data[0].Title,
			Subtitle:      reply.Data[0].Subtitle,
			Img:           reply.Data[0].Img,
			Price:         reply.Data[0].Price,
			UserName:      userNameSrc[rand.Intn(len(userNameSrc))],
		}
		_, e := orderClient.CreateOrder(ctx, &model)
		e1 := utils.CheckRPCError(e)
		utils.CheckErr(e1, "")
	}
	models.DB.Close()
}

func main() {
	conf.InitConfig()
	// 定义一个cron运行器
	loc, e := time.LoadLocation("Asia/Shanghai")
	utils.Fatalf(e, "")
	c := cron.NewWithLocation(loc)
	// 每分钟执行一次
	err := c.AddFunc(fmt.Sprintf("0 */1 * * *"), run)
	utils.Fatalf(err, "")
	c.Start()
	c1 := make(chan struct{})
	<-c1
}
