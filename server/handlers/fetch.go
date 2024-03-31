package handlers

import (
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/mathutil"
	"github.com/kmou424/resonance-dataserver/cache"
	"github.com/kmou424/resonance-dataserver/database/model"
	"github.com/kmou424/resonance-dataserver/database/repositories"
	"github.com/kmou424/resonance-dataserver/internal/kits/strkit"
	"github.com/kmou424/resonance-dataserver/pojo"
	"net/http"
	"time"
)

var GetGoodsInfo gin.HandlerFunc = func(c *gin.Context) {
	station := c.Query("station")
	goodName := c.Query("good")

	var goods []pojo.Good

	goodsMapper := repositories.GoodsMapper.Find(goodName, station)
	for _, mapper := range goodsMapper {
		goods = append(goods, *mapperToGood(&mapper))
	}

	c.JSON(http.StatusOK, goods)
}

func walkGoodExtraByID(id string, walk func(pojo.GoodExtra)) {
	keys := []string{
		strkit.Concat(id, "_update_time"),
		strkit.Concat(id, "_price"),
		strkit.Concat(id, "_next_trend"),
	}
	vals, _ := cache.MGet(keys...)

	getValueOrDef := func(idx int, def any) any {
		if vals[idx] == nil {
			log.Warn("cached value is nil", "key", keys[idx])
			return def
		}
		return vals[idx]
	}

	var goodExtra pojo.GoodExtra

	// update_time
	{
		updateTimestamp := getValueOrDef(0, 0)
		updateTimestampInt64 := mathutil.MustInt64(updateTimestamp)

		goodExtra.UpdateTimestamp = updateTimestampInt64
		goodExtra.UpdateTime = time.Unix(updateTimestampInt64, 0).Format(time.DateTime)
	}

	// price
	{
		price := getValueOrDef(1, -1)
		goodExtra.Price = mathutil.MustInt(price)
	}

	// next_trend
	{
		nextTrend := getValueOrDef(2, 0)
		goodExtra.NextTrend = mathutil.MustInt(nextTrend)
	}

	walk(goodExtra)
}

func mapperToGood(mapper *model.GoodsMapper) (good *pojo.Good) {
	mapper.GoodBase.ID = mapper.ID
	walkGoodExtraByID(mapper.ID, func(extra pojo.GoodExtra) {
		good = &pojo.Good{
			GoodBase:  mapper.GoodBase,
			GoodExtra: extra,
		}
	})
	return
}
