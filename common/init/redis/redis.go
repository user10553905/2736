package redis

import (
	"common/nacos"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     nacos.Config.Redis.Addr,
		Password: nacos.Config.Redis.Pass,
		DB:       nacos.Config.Redis.Db,
	})
	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis连接成功")
}
