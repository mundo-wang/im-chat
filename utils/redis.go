package utils

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
)

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", Config.Redis.IP, Config.Redis.Port),
		Password:     Config.Redis.Password,
		DB:           Config.Redis.DB,
		PoolSize:     Config.Redis.PoolSize,
		MinIdleConns: Config.Redis.MinIdleConn,
	})
}
