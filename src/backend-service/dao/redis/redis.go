package redis

import (
	"backend-service/settings"
	"fmt"
	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var RDB *redis.Client

// Init 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	RDB = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
		PoolSize: cfg.PoolSize,
	})

	_, err = RDB.Ping().Result()
	return
}

func Close() {
	_ = RDB.Close()
}
