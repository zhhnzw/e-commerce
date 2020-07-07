package models

import (
	"backend-service/conf"
	"backend-service/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB
var DBForFrontService *gorm.DB

func InitGorm() {
	sourceURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/backend-service?parseTime=true&loc=Local",
		conf.Config.Mysql.User,
		conf.Config.Mysql.Password,
		conf.Config.Mysql.Host,
		conf.Config.Mysql.Port)
	log.Println("mysql init:" + sourceURL)
	db, err := gorm.Open("mysql", sourceURL)
	db.LogMode(true)
	if err != nil {
		utils.Logf(err, "连接数据库失败")
		panic("")
	}
	DB = db
}

func InitGormForFrontService() {
	sourceURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/front-service?parseTime=true&loc=Local",
		conf.Config.Mysql.User,
		conf.Config.Mysql.Password,
		conf.Config.Mysql.Host,
		conf.Config.Mysql.Port)
	log.Println("mysql init:" + sourceURL)
	db, err := gorm.Open("mysql", sourceURL)
	db.LogMode(true)
	if err != nil {
		utils.Logf(err, "连接数据库失败")
		panic("")
	}
	DBForFrontService = db
}
