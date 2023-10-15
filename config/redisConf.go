package config

import (
	"fmt"
	"github.com/go-redis/redis"
)

var (
	redisClient *redis.Client
)

func NewRedisClient() *redis.Client {
	var redisClient *redis.Client
	redisClient = redis.NewClient(&redis.Options{
		//Addr:     config.C.Redis.Addr,
		//Password: config.C.Redis.Password,
		//DB:       config.C.Redis.DB,
		Addr:     "localhost:6379",
		Password: "111111",
		DB:       0,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("redis failed connected:%s", err.Error()))
	}
	return redisClient
}
