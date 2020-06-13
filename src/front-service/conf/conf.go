package conf

import (
	"flag"
	"github.com/jinzhu/configor"
	"log"
	"os"
)

var Config = struct {
	RunMode          string
	AppPort          string
	GoodsServiceAddr string

	Mysql struct {
		Host     string
		Port     uint   `default:"3306"`
		User     string `default:"root"`
		Password string `env:"DBPassword"`
	}

	Redis struct {
		Host     string
		Port     uint `default:"6379"`
		Password string
		Db       int
	}

	Dev struct {
		Mysql struct {
			Host     string
			Port     uint   `default:"3306"`
			User     string `default:"root"`
			Password string `required:"true"`
		}
		Redis struct {
			Host     string
			Port     uint `default:"6379"`
			Password string
			Db       int
		}
	}

	Prev struct {
		Mysql struct {
			Host     string
			Port     uint   `default:"3306"`
			User     string `default:"root"`
			Password string `required:"true"`
		}
		Redis struct {
			Host     string
			Port     uint `default:"6379"`
			Password string
			Db       int
		}
	}

	Prod struct {
		Mysql struct {
			Host     string
			Port     uint   `default:"3306"`
			User     string `default:"root"`
			Password string `required:"true"`
		}
		Redis struct {
			Host     string
			Port     uint `default:"6379"`
			Password string
			Db       int
		}
	}
}{}

func InitConfig() {
	readRunParam()
	var configDir string
	workDir := os.Getenv("WORK_DIR")
	if len(workDir) > 0 {
		configDir = workDir
	}
	if err := configor.Load(&Config, configDir+"conf/config.yaml"); err != nil {
		panic(err)
	}
	switch Config.RunMode {
	case "dev":
		Config.Mysql.Host = Config.Dev.Mysql.Host
		Config.Mysql.Port = Config.Dev.Mysql.Port
		Config.Mysql.User = Config.Dev.Mysql.User
		Config.Mysql.Password = Config.Dev.Mysql.Password
		Config.Redis.Host = Config.Dev.Redis.Host
		Config.Redis.Port = Config.Dev.Redis.Port
		Config.Redis.Password = Config.Dev.Redis.Password
		Config.Redis.Db = Config.Dev.Redis.Db
	case "prev":
		Config.Mysql.Host = Config.Prev.Mysql.Host
		Config.Mysql.Port = Config.Prev.Mysql.Port
		Config.Mysql.User = Config.Prev.Mysql.User
		Config.Mysql.Password = Config.Prev.Mysql.Password
		Config.Redis.Host = Config.Prev.Redis.Host
		Config.Redis.Port = Config.Prev.Redis.Port
		Config.Redis.Password = Config.Prev.Redis.Password
		Config.Redis.Db = Config.Prev.Redis.Db
	case "prod":
		Config.Mysql.Host = Config.Prod.Mysql.Host
		Config.Mysql.Port = Config.Prod.Mysql.Port
		Config.Mysql.User = Config.Prod.Mysql.User
		Config.Mysql.Password = Config.Prod.Mysql.Password
		Config.Redis.Host = Config.Prod.Redis.Host
		Config.Redis.Port = Config.Prod.Redis.Port
		Config.Redis.Password = Config.Prod.Redis.Password
		Config.Redis.Db = Config.Prod.Redis.Db
	default:
		log.Fatal("配置文件: utils/config.yaml runmode配置有误，可配置选择：dev、prev、prod")
	}
	log.Println("run mode:", Config.RunMode)
}

func readRunParam() {
	var workDir string
	flag.StringVar(&workDir, "WORK_DIR", "", "工作目录")
	flag.Parse()
	if len(workDir) > 0 {
		log.Printf("WORK_DIR:%s", workDir)
		os.Setenv("WORK_DIR", workDir)
	}
}
