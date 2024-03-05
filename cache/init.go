package cache

import (
	"github.com/gookit/goutil/envutil"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     envutil.Getenv("REDIS_ADDR", "localhost:6379"),
		Password: envutil.Getenv("REDIS_PWD", ""),
		DB:       0,
		PoolSize: 20,
	})
}
