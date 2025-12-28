package repository

import "github.com/redis/go-redis/v9"

var RedisCilent redis.UniversalClient

func InitRedis() {
	// RedisCilent = redis.NewClient(&redis.Options{
	// 	Addr: "172.23.84.152:6379", //后续从配置文件读取
	// 	// Password: "", // no password set
	// 	// DB:       0,  // use default DB
	// })
	RedisCilent = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{"172.23.84.152:6379"},
	})
}
