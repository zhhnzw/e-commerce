package test

import (
	"backend-service/conf"
	"backend-service/routers"
	"backend-service/utils"
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
)

var Router *gin.Engine

func InitTestServer() {
	// 修改工作目录以正常读取配置文件启动服务
	if err := os.Setenv("WORK_DIR", "../../"); err != nil {
		utils.Fatalf(err, "环境变量设置失败")
	}
	if Router == nil {
		conf.InitConfig()
		gin.SetMode(gin.DebugMode)
		Router = routers.InitRouter()
	}
}

func SendRequest(URL, param, method string) (body []byte, statusCode int) {
	InitTestServer()
	req, err := http.NewRequest(method, URL, bytes.NewBuffer([]byte(param)))
	utils.Fatalf(err, "")
	req.Header.Set("Content-Type", "application/json")
	Recorder := httptest.NewRecorder()
	Router.ServeHTTP(Recorder, req)
	body, err = ioutil.ReadAll(Recorder.Result().Body)
	utils.Fatalf(err, "")
	defer Recorder.Result().Body.Close()
	return body, Recorder.Code
}
