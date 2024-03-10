package rediskit

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/gookit/goutil/mathutil"
	"github.com/kmou424/resonance-dataserver/cache"
	"github.com/redis/go-redis/v9"
)

func GetValueFromRedis(key string, def any) any {
	val, err := cache.Get(key)
	if errors.Is(err, redis.Nil) {
		log.Warn(fmt.Sprintf("get value [%s] failed: not exists", key))
		return def
	}
	if err != nil {
		log.Warn(fmt.Sprintf("get value [%s] failed", key))
		return def
	}
	return val
}

func GetIntFromRedis(key string, def int) int {
	value := GetValueFromRedis(key, def)
	return mathutil.MustInt(value)
}
