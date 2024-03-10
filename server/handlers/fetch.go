package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/mathutil"
	model2 "github.com/kmou424/resonance-dataserver/database/model"
	"github.com/kmou424/resonance-dataserver/database/repositories"
	"github.com/kmou424/resonance-dataserver/internal/kits/hashkit"
	"github.com/kmou424/resonance-dataserver/internal/kits/rediskit"
	"github.com/kmou424/resonance-dataserver/internal/kits/strkit"
	"github.com/kmou424/resonance-dataserver/pojo"
	"github.com/kmou424/resonance-dataserver/server/errors"
	"net/http"
	"strings"
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

var GetFullGoodsInfo gin.HandlerFunc = func(c *gin.Context) {
	station := c.Query("station")
	if station == "" {
		panic(errors.BadRequest(`you must provide "station" to query goods`))
	}

	show := c.Query("show")
	switch show {
	case "full":
	case "unknown":
		break
	default:
		show = "full"
	}

	existGoodsMapper := repositories.GoodsMapper.Find("", station)
	existGoodsMap := make(map[string]model2.GoodsMapper)
	for _, mapper := range existGoodsMapper {
		existGoodsMap[mapper.ID] = mapper
	}

	var fullGoods []pojo.FullGood
	value := rediskit.GetValueFromRedis(strkit.Concat(hashkit.MD5(station), "_goods_list"), "")
	if value != "" {
		goodsIdList := mathutil.MustString(value)
		for _, goodId := range strings.Split(goodsIdList, ",") {
			var mapper model2.GoodsMapper
			if goodMapper, ok := existGoodsMap[goodId]; ok {
				// filter known goods
				if show == "unknown" {
					continue
				}
				mapper = goodMapper
			} else {
				mapper = model2.GoodsMapper{
					ID:      goodId,
					Name:    "Unknown",
					Station: station,
				}
			}
			fullGood := pojo.FullGood{
				Id:   goodId,
				Good: *mapperToGood(&mapper),
			}
			fullGoods = append(fullGoods, fullGood)
		}
	}

	c.JSON(http.StatusOK, fullGoods)
}

func mapperToGood(mapper *model2.GoodsMapper) *pojo.Good {
	goodId := mapper.ID
	updateTimestamp := rediskit.GetValueFromRedis(strkit.Concat(goodId, "_update_time"), time.Now().Unix()+1000000)
	updateTimestampInt := mathutil.MustInt64(updateTimestamp)
	good := &pojo.Good{
		Name:            mapper.Name,
		Station:         mapper.Station,
		Price:           rediskit.GetIntFromRedis(strkit.Concat(goodId, "_price"), -1),
		NextTrend:       rediskit.GetIntFromRedis(strkit.Concat(goodId, "_next_trend"), 0),
		UpdateTime:      time.Unix(updateTimestampInt, 0).Format(time.DateTime),
		UpdateTimestamp: updateTimestampInt,
	}
	return good
}
