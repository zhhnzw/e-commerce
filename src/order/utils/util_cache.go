package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"order/conf"
	"time"
)

var RedisClient *redis.Client

func InitRedis() {
	addr := fmt.Sprintf("%s:%d", conf.Config.Redis.Host, conf.Config.Redis.Port)
	log.Println("redis init:" + addr)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: conf.Config.Redis.Password,
		DB:       conf.Config.Redis.Db,
	})
	RedisClient = client
	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}
}

type Cache struct {
	RedisKeyName string
	RedisKeyType string
	Result       interface{} //  缓存结果
}

func (cache *Cache) StoreStringCache() error {
	return RedisClient.Set(cache.RedisKeyName, ToJson(cache.Result), time.Minute).Err() // 缓存1分钟
}

func (cache *Cache) GetStringCache() (string, error) {
	return RedisClient.Get(cache.RedisKeyName).Result()
}
