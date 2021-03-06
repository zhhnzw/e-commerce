package v1

import (
	"backend-service/pb"
	"backend-service/settings"
	"backend-service/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

const (
	StatusOk       string = "操作成功!"
	GetFailed      string = "操作失败!"
	CreateFailed   string = "创建失败!"
	CreateRepeated string = "重复创建!"
	DeleteFailed   string = "删除失败!"
	UpdateFailed   string = "更新失败!"
)

var goodsClient pb.GoodsClient

func InitGoodsRPCClient(cfg *settings.GoodsConfig) (err error) {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	zap.L().Info("goods gRPC service connect:" + addr)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	goodsClient = pb.NewGoodsClient(conn)
	return
}

type Goods struct {
	Id            int            `json:"-"`
	GoodsUuid     string         `json:"goodsUuid" form:"goodsUuid"`
	GoodsFrom     string         `json:"goodsFrom" form:"goodsFrom"`
	GoodsTypeId   int64          `json:"goodsTypeId" gorm:"goodsTypeId"`
	PrimaryType   string         `json:"primaryType" form:"primaryType"`
	SecondaryType string         `json:"secondaryType" form:"secondaryType"`
	Img           string         `json:"img"`
	Imgs          string         `json:"imgs"`
	IsValid       bool           `json:"isValid" gorm:"default:true,column:is_valid" form:"isValid"`
	Title         string         `json:"title"`
	SubTitle      string         `json:"subTitle"`
	Price         int            `json:"price"`
	PublishDate   utils.JSONTime `json:"publishDate" form:"-" gorm:"-"`
	CreatedTime   time.Time      `json:"-" form:"-" gorm:"-"`
	UpdatedTime   utils.JSONTime `json:"-" form:"-" gorm:"-"`
	DeletedTime   time.Time      `json:"-" form:"-" gorm:"-"`
	PageSize      int64          `gorm:"-" json:"-" form:"pageSize"`
	PageIndex     int64          `gorm:"-" json:"-" form:"pageIndex"`
}

// 查询商品列表
func GetGoodsList(c *gin.Context) {
	var model Goods
	err := c.ShouldBindQuery(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Resp{Message: "参数错误", Code: "1"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	request := pb.GoodsRequest{
		GoodsUuid:   model.GoodsUuid,
		GoodsTypeId: model.GoodsTypeId,
		PageIndex:   model.PageIndex,
		PageSize:    model.PageSize,
	}
	reply, err := goodsClient.GetGoodsList(ctx, &request)
	err = utils.CheckRPCError(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Resp{Message: err.Error(), Code: "1"})
	} else {
		c.JSON(http.StatusOK, utils.Resp{Data: reply, Code: "0", Message: StatusOk})
	}
}

func GetGoods(c *gin.Context) {
	uuid := c.Param("uuid")
	if len(uuid) == 0 {
		c.JSON(http.StatusBadRequest, utils.Resp{Message: "参数错误", Code: "1"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	request := pb.GoodsRequest{
		GoodsUuid: uuid,
	}
	reply, err := goodsClient.GetGoodsDetail(ctx, &request)
	err = utils.CheckRPCError(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Resp{Message: err.Error(), Code: "1"})
	} else {
		c.JSON(http.StatusOK, utils.Resp{Data: reply, Code: "0", Message: StatusOk})
	}
}

func GetHotGoodsList(c *gin.Context) {
	var model Goods
	err := c.ShouldBindQuery(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Resp{Message: "参数错误", Code: "1"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	request := pb.GoodsRequest{
		PrimaryType:   model.PrimaryType,
		SecondaryType: model.SecondaryType,
		PageIndex:     model.PageIndex,
		PageSize:      model.PageSize,
	}
	reply, err := goodsClient.GetGoodsHotList(ctx, &request)
	err = utils.CheckRPCError(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Resp{Message: err.Error(), Code: "1"})
	} else {
		c.JSON(http.StatusOK, utils.Resp{Data: reply, Code: "0", Message: StatusOk})
	}
}

func GetGoodsStatistic(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	request := pb.GoodsRequest{}
	reply, err := goodsClient.GetGoodsStatistic(ctx, &request)
	err = utils.CheckRPCError(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Resp{Message: err.Error(), Code: "1"})
	} else {
		c.JSON(http.StatusOK, utils.Resp{Data: reply, Code: "0", Message: StatusOk})
	}
}
