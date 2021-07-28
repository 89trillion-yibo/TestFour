package model

import (
	"fmt"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

//连接redis数据库
func RedisCli(httpport string) (error)  {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     httpport,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		fmt.Println("connect redis failed")
		return  err
	}
	fmt.Println("RedisClient:", RedisClient)
	return nil
}
