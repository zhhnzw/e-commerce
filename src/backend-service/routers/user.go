package routers

import (
	"backend-service/controller/v1"
	"github.com/gin-gonic/gin"
)

func SetUserRouter(router *gin.Engine) {
	router.GET("/v1/statistic/user", v1.GetStatisticForUser)
	router.POST("/v1/logout", v1.Logout)
	router.POST("/v1/alterPwd", v1.AlterPwd)
	userRouter := router.Group("/v1/sys/user")
	userRouter.GET("", v1.GetUsers)
	userRouter.POST("", v1.CreateUser)
}
