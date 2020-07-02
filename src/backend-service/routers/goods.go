package routers

import (
	"backend-service/controller/v1"
	"github.com/gin-gonic/gin"
)

func SetGoodsRouter(router *gin.Engine) {
	router.GET("/v1/hot/goods", v1.GetHotGoodsList)
	router.GET("/v1/statistic/goods", v1.GetGoodsStatistic)
	goodsRouter := router.Group("/v1/goods")
	goodsRouter.GET("", v1.GetGoodsList)
	goodsRouter.GET("/:uuid", v1.GetGoods)
}
