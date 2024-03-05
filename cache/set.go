package cache

import (
	"github.com/charmbracelet/log"
	"time"
)

func Set(key string, value any, duration time.Duration) error {
	ctx, _ := contextTimeout()
	err := redisClient.Set(ctx, key, value, duration).Err()
	if err != nil {
		log.Error("caching value failed", "key", key, "error", err)
	}
	return err
}
