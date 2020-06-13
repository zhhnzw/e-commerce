package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"goods/conf"
	"time"
)

var RedisClient *redis.Client

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Config.Redis.Host, conf.Config.Redis.Port),
		Password: conf.Config.Redis.Password,
		DB:       conf.Config.Redis.Db,
	})
	RedisClient = client
	if _, err := client.Ping().Result(); err != nil {
		panic("redis 连接异常")
	}
}

type Cache struct {
	RedisKeyName string
	RedisKeyType string
	Result       string //  缓存结果
}

func (cache *Cache) StoreStringCache() error {
	return RedisClient.Set(cache.RedisKeyName, cache.Result, time.Minute).Err() // 缓存1分钟
}

func (cache *Cache) GetStringCache() (string, error) {
	return RedisClient.Get(cache.RedisKeyName).Result()
}
