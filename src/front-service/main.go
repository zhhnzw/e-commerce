package main

import (
	"front-service/conf"
	"front-service/models"
	"front-service/routers"
	"front-service/utils"
	"github.com/gin-gonic/gin"
)

func runServer() {
	conf.InitConfig()
	models.InitGorm()
	utils.InitRedis()
	gin.SetMode(gin.DebugMode)
	router := routers.InitRouter()
	port := conf.Config.AppPort
	router.Run(":" + port)
}

func main() {
	runServer()
}
