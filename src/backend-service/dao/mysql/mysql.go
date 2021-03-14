package mysql

import (
	"backend-service/settings"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var DB *gorm.DB
var DBForFrontService *gorm.DB

func InitGorm(cfg *settings.MySQLConfig) (err error) {
	sourceURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/backend-service?parseTime=true&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port)
	zap.L().Info("mysql connect:" + sourceURL)
	DB, err = gorm.Open("mysql", sourceURL)
	DB.LogMode(true)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	DB.DB().SetMaxOpenConns(cfg.MaxOpenConns)
	DB.DB().SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

func InitGormForFrontService(cfg *settings.MySQLConfig) (err error) {
	sourceURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/front-service?parseTime=true&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port)
	zap.L().Info("mysql connect:" + sourceURL)
	DBForFrontService, err = gorm.Open("mysql", sourceURL)
	DBForFrontService.LogMode(true)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	DBForFrontService.DB().SetMaxOpenConns(cfg.MaxOpenConns)
	DBForFrontService.DB().SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

func Close() {
	_ = DB.Close()
	_ = DBForFrontService.Close()
}
