package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/mathutil"
	"github.com/kmou424/resonance-dataserver/database/model"
	"github.com/kmou424/resonance-dataserver/database/repositories"
	"github.com/kmou424/resonance-dataserver/internal/kits/rediskit"
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

func mapperToGood(mapper *model.GoodsMapper) *pojo.Good {
	goodId := mapper.ID
	updateTimestamp := rediskit.GetValueFromRedis(strkit.Concat(goodId, "_update_time"), time.Now().Unix()+1000000)
	updateTimestampInt := mathutil.MustInt64(updateTimestamp)
	good := &pojo.Good{
		GoodBase: mapper.GoodBase,
		GoodExtra: pojo.GoodExtra{
			Price:           rediskit.GetIntFromRedis(strkit.Concat(goodId, "_price"), -1),
			NextTrend:       rediskit.GetIntFromRedis(strkit.Concat(goodId, "_next_trend"), 0),
			UpdateTime:      time.Unix(updateTimestampInt, 0).Format(time.DateTime),
			UpdateTimestamp: updateTimestampInt,
		},
	}
	return good
}
