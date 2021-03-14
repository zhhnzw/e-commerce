package routers

import (
	"backend-service/controller"
	"backend-service/controller/v1"
	"backend-service/logger"
	"backend-service/settings"
	"backend-service/utils"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup() *gin.Engine {
	if settings.Conf.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(logger.GinLogger(), logger.GinRecovery(true))
	config := cors.DefaultConfig()
	config.AllowOrigins = settings.Conf.AllowOrigins
	config.AllowCredentials = true
	router.Use(gin.Recovery())
	address := fmt.Sprintf("%s:%d", settings.Conf.RedisConfig.Host, settings.Conf.RedisConfig.Port)
	store, err := redis.NewStore(16, "tcp", address, settings.Conf.RedisConfig.Password, []byte("secret"))
	utils.CheckErr(err, "")
	router.Use(sessions.Sessions("session", store))
	router.POST("/v1/login", v1.Login)
	router.Use(controller.SetAuthMiddleware())
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
