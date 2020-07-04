package v1

import (
	"backend-service/conf"
	"backend-service/pb"
	"backend-service/utils"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"time"
)

var orderClient pb.OrderClient

func InitOrderRPCClient() {
	conn, err := grpc.Dial(conf.Config.OrderServiceAddr, grpc.WithInsecure())
	utils.Fatalf(err, "")
	orderClient = pb.NewOrderClient(conn)
}

type Order struct {
	Id            int
	GoodsUuid     string         `json:"goodsUuid" form:"goodsUuid"`
	GoodsTypeId   int64          `json:"goodsTypeId" form:"goodsTypeId"`
	PrimaryType   string         `json:"primaryType" form:"primaryType"`
	SecondaryType string         `json:"secondaryType" form:"secondaryType"`
	Img           string         `json:"img" form:"img"`
	Title         string         `json:"title" form:"title"`
	Subtitle      string         `json:"subtitle" form:"subtitle"`
	Price         int64          `json:"price" form:"price"`
	OrderStatus   string         `json:"orderStatus" form:"orderStatus"`
	CreatedTime   time.Time      `json:"-" form:"-" gorm:"-"`
	UpdatedTime   utils.JSONTime `json:"-" form:"-" gorm:"-"`
	DeletedTime   time.Time      `json:"-" form:"-" gorm:"-"`
	PageSize      int64          `gorm:"-" json:"-" form:"pageSize"`
	PageIndex     int64          `gorm:"-" json:"-" form:"pageIndex"`
}

func CreateOrder(c *gin.Context) {
	var model Order
	err := c.ShouldBindJSON(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Resp{Message: "参数错误:%s" + err.Error(), Code: "1"})
		return
	}
	log.Printf("%+v", model)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	request := pb.OrderRequest{
		GoodsUuid:     model.GoodsUuid,
		GoodsTypeId:   model.GoodsTypeId,
		PrimaryType:   model.PrimaryType,
		SecondaryType: model.SecondaryType,
		Img:           model.Img,
		Title:         model.Title,
		Subtitle:      model.Subtitle,
		Price:         model.Price,
		OrderStatus:   model.OrderStatus,
	}
	reply, err := orderClient.CreateOrder(ctx, &request)
	err = utils.CheckRPCError(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Resp{Message: err.Error(), Code: "1"})
	} else {
		c.JSON(http.StatusOK, utils.Resp{Data: reply, Code: "0", Message: StatusOk})
	}
}

func GetOrderList(c *gin.Context) {
	var model Order
	err := c.ShouldBindQuery(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Resp{Message: "参数错误", Code: "1"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	request := pb.OrderRequest{
		GoodsUuid:     model.GoodsUuid,
		PrimaryType:   model.PrimaryType,
		SecondaryType: model.SecondaryType,
		PageIndex:     model.PageIndex,
		PageSize:      model.PageSize,
	}
	reply, err := orderClient.GetOrderList(ctx, &request)
	err = utils.CheckRPCError(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Resp{Message: err.Error(), Code: "1"})
	} else {
		c.JSON(http.StatusOK, utils.Resp{Data: reply, Code: "0", Message: StatusOk})
	}
}

func GetOrderStatistic(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	request := pb.OrderRequest{}
	reply, err := orderClient.GetOrderStatistic(ctx, &request)
	err = utils.CheckRPCError(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Resp{Message: err.Error(), Code: "1"})
	} else {
		c.JSON(http.StatusOK, utils.Resp{Data: reply, Code: "0", Message: StatusOk})
	}
}
