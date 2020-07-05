package routers

import (
	"front-service/controller/v1"
	"github.com/gin-gonic/gin"
)

func SetUserRouter(router *gin.Engine) {
	router.POST("/v1/logout", v1.Logout)
	router.POST("/v1/alterPwd", v1.AlterPwd)
	router.GET("/v1/user", v1.GetUsers)
}
