package main

import (
	"front-service/conf"
	"front-service/controller/v1"
	"front-service/routers"
	"github.com/gin-gonic/gin"
)

func runServer() {
	conf.InitConfig()
	gin.SetMode(gin.DebugMode)
	router := routers.InitRouter()
	v1.InitGoodsRPCClient()
	port := conf.Config.AppPort
	router.Run(":" + port)
}

func main() {
	runServer()
}
