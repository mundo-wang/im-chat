package utils

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"im-chat/conf"
	"sync"
)

var (
	redisClient *redis.Client
	redisOnce   sync.Once
)

func NewRedisClient() *redis.Client {
	redisOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:         fmt.Sprintf("%s:%d", conf.Config.Redis.IP, conf.Config.Redis.Port),
			Password:     conf.Config.Redis.Password,
			DB:           conf.Config.Redis.DB,
			PoolSize:     conf.Config.Redis.PoolSize,
			MinIdleConns: conf.Config.Redis.MinIdleConn,
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
