package utils

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
)

var (
	redisClient *redis.Client
	redisOnce   sync.Once
)

func NewRedisClient() *redis.Client {
	redisOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:         fmt.Sprintf("%s:%d", Config.Redis.IP, Config.Redis.Port),
			Password:     Config.Redis.Password,
			DB:           Config.Redis.DB,
			PoolSize:     Config.Redis.PoolSize,
			MinIdleConns: Config.Redis.MinIdleConn,
		})
	})
	return redisClient
}

func CloseRedis() error {
	if redisClient != nil {
		return redisClient.Close()
	}
	return nil
}
