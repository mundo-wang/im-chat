package utils

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/mundo-wang/wtool/wlog"
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
	wlog.Info("InitRedis complete!").Log()
}
