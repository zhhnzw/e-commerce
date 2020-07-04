package routers

import (
	"backend-service/controller/v1"
	"github.com/gin-gonic/gin"
)

func SetOrderRouter(router *gin.Engine) {
	router.GET("/v1/statistic/order", v1.GetOrderStatistic)
	goodsRouter := router.Group("/v1/order")
	goodsRouter.GET("", v1.GetOrderList)
	goodsRouter.POST("", v1.CreateOrder)
}
