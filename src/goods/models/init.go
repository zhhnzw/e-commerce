package models

import (
	"database/sql/driver"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"goods/conf"
	"goods/utils"
	"time"
)

var DB *gorm.DB

func InitGorm() {
	sourceURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/shihuo?parseTime=true&loc=Local",
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

// 用于gorm的time类型格式化  http://axiaoxin.com/article/241/
type JSONTime struct {
	time.Time
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// 可用于json自定义格式的反序列化
func (t JSONTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	var err error
	t.Time, err = time.ParseInLocation(`"2006-01-02 15:04:05"`, string(b), time.Local)
	return err
}

func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
