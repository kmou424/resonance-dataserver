package cache

import (
	"errors"
	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"
)

func MGet(keys ...string) (vals []any, err error) {
	ctx, _ := contextTimeout()
	vals, err = redisClient.MGet(ctx, keys...).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		log.Error("get values from redis failed", "keys", keys, "error", err)
	}
	return
}
