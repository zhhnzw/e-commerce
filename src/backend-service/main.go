package main

import (
	"backend-service/conf"
	"backend-service/controller/v1"
	"backend-service/models"
	"backend-service/routers"
	"backend-service/utils"
	"github.com/gin-gonic/gin"
)

func runServer() {
	conf.InitConfig()
	models.InitGorm()
	utils.InitRedis()
	gin.SetMode(gin.DebugMode)
	router := routers.InitRouter()
	v1.InitGoodsRPCClient()
	port := conf.Config.AppPort
	router.Run(":" + port)
}

func main() {
	runServer()
}
