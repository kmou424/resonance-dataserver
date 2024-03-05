package handlers

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/mathutil"
	"github.com/kmou424/resonance-dataserver/cache"
	"github.com/kmou424/resonance-dataserver/database/repositories"
	"github.com/kmou424/resonance-dataserver/internal/kits/strkit"
	"github.com/kmou424/resonance-dataserver/model"
	"github.com/kmou424/resonance-dataserver/server/errors"
	"net/http"
	"time"
)

var GetGoodsInfo gin.HandlerFunc = func(c *gin.Context) {
	station := c.Query("station")
	goodName := c.Query("good")
	if station == "" && goodName == "" {
		panic(errors.BadRequest("at least provide an argument: station or good"))
	}

	getValueFromRedis := func(key string, def any) any {
		val, err := cache.Get(key)
		if err != nil {
			log.Warn(fmt.Sprintf("get entry of good info failed: [%s]", key))
			return def
		}
		return val
	}

	getIntFromRedis := func(key string, def int) int {
		value := getValueFromRedis(key, def)
		return mathutil.MustInt(value)
	}

	var goods []*model.Good

	goodsMapper := repositories.GoodsMapper.Find(goodName, station)
	for _, mapper := range goodsMapper {
		goodId := mapper.ID
		updateTimestamp := getValueFromRedis(strkit.Concat(goodId, "_update_time"), time.Now().Unix()+1000000)
		updateTimestampInt := mathutil.MustInt64(updateTimestamp)
		good := &model.Good{
			Name:            mapper.Name,
			Station:         mapper.Station,
			Price:           getIntFromRedis(strkit.Concat(goodId, "_price"), -1),
			NextTrend:       getIntFromRedis(strkit.Concat(goodId, "_next_trend"), 0),
			UpdateTime:      time.Unix(updateTimestampInt, 0).Format(time.DateTime),
			UpdateTimestamp: updateTimestampInt,
		}
		goods = append(goods, good)
	}

	c.JSON(http.StatusOK, goods)
}
