package utils

import (
	"context"
	"errors"
	"github.com/go-redis/redis"
	"goods/pb"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

// 拦截器和rpc处理函数传值用的
var InterceptorKey = make(map[string]string)

// 缓存key的前缀
var CacheKeyPrefix = "goods_cache_"

func CacheInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("before cacheInterceptor. req: %+v", req)
	if _, ok := req.(*pb.GoodsRequest); ok {
		visitKey := "goods_visit_" + info.FullMethod
		res := RedisClient.Get(visitKey)
		if r, e := res.Result(); e != nil {
			Logf(e, "redis get visitKey failed, keyName:"+visitKey)
			//终止后续操作，不再进入controller
			return nil, e
		} else {
			// 当访问量大于1次，需要操作缓存
			if visitNum, err := strconv.Atoi(r); err != nil {
				Logf(err, "redis key string to int failed, keyName:"+visitKey)
				return nil, e

			} else if visitNum == 1 {
				resp, err := handler(ctx, req)
				return resp, err
			} else {
				cache := Cache{
					RedisKeyName: CacheKeyPrefix + info.FullMethod,
					Result:       "",
				}
				// 若缓存存在
				if resp, err := cache.GetStringCache(); err == nil {
					// 把cache变量传过去
					InterceptorKey[cache.RedisKeyName] = resp
					resp, err := handler(ctx, req)
					return resp, err

					// 若缓存不存在
				} else if err == redis.Nil {
					// 在处理完请求后把结果缓存下来
					resp, err := handler(ctx, req)
					CheckErr(err, "%s处理出现异常", info.FullMethod)
					cache.Result = ToJson(resp)
					err1 := cache.StoreStringCache()
					CheckErr(err1, "缓存失败, 缓存数据:%+v", resp)
					return resp, err
				} else {
					Logf(err, "redis异常!!!")
					return nil, errors.New("")
				}
			}
		}
	}
	log.Printf("after cacheInterceptor")
	return nil, nil
}
