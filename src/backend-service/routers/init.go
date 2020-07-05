package routers

import (
	"backend-service/conf"
	"backend-service/controller/v1"
	"backend-service/utils"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func InitRouter() *gin.Engine {

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout) //gin日志
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = conf.Config.AllowOrigins
	config.AllowCredentials = true
	router.Use(cors.New(config))
	//router.Use(gin.Logger())
	router.Use(gin.Recovery())
	address := fmt.Sprintf("%s:%d", conf.Config.Redis.Host, conf.Config.Redis.Port)
	store, err := redis.NewStore(16, "tcp", address, conf.Config.Redis.Password, []byte("secret"))
	utils.Fatalf(err, "")
	router.Use(sessions.Sessions("session", store))
	router.POST("/v1/login", v1.Login)
	router.Use(utils.SetAuthMiddleware())
	SetUserRouter(router)
	SetGoodsRouter(router)
	SetOrderRouter(router)
	router.NoMethod(func(c *gin.Context) {
		c.JSON(
			http.StatusMethodNotAllowed,
			utils.Resp{Message: "方法不允许"})
	})
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound,
			utils.Resp{Message: "资料不存在"})
	})
	return router
}
