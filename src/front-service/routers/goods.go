package routers

import (
	"front-service/controller/v1"
	"github.com/gin-gonic/gin"
)

func SetGoodsRouter(router *gin.Engine) {
	router.GET("/v1/hot", v1.GetHotGoodsList)
	goodsRouter := router.Group("/v1/goods")
	goodsRouter.GET("", v1.GetGoodsList)
	goodsRouter.GET("/:uuid", v1.GetGoods)
}
