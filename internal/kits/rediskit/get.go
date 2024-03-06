package rediskit

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/gookit/goutil/mathutil"
	"github.com/kmou424/resonance-dataserver/cache"
)

func GetValueFromRedis(key string, def any) any {
	val, err := cache.Get(key)
	if err != nil {
		log.Warn(fmt.Sprintf("get entry of good info failed: [%s]", key))
		return def
	}
	return val
}

func GetIntFromRedis(key string, def int) int {
	value := GetValueFromRedis(key, def)
	return mathutil.MustInt(value)
}
