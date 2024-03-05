package cache

import (
	"errors"
	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"
)

func Get(key string) (val any, err error) {
	ctx, _ := contextTimeout()
	val, err = redisClient.Get(ctx, key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		log.Error("get value from redis failed", "key", key, "error", err)
	}
	return
}
