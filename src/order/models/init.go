package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"order/conf"
	"order/utils"
)

var DB *gorm.DB

func InitGorm() {
	sourceURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/order?parseTime=true&loc=Local",
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
