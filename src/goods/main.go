package main

import (
	"github.com/gin-gonic/gin"
	"goods/controller/v1"
	"goods/conf"
	"goods/models"
	"goods/utils"
)

func runServer() {
	conf.InitConfig()
	utils.InitRedis()
	models.InitGorm()
	gin.SetMode(gin.DebugMode)
	v1.Run()
}

func main() {
	runServer()
}
