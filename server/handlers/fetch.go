package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/mathutil"
	"github.com/kmou424/resonance-dataserver/database/repositories"
	"github.com/kmou424/resonance-dataserver/internal/kits/hashkit"
	"github.com/kmou424/resonance-dataserver/internal/kits/rediskit"
	"github.com/kmou424/resonance-dataserver/internal/kits/strkit"
	"github.com/kmou424/resonance-dataserver/model"
	"github.com/kmou424/resonance-dataserver/server/errors"
	"net/http"
	"strings"
	"time"
)

var GetGoodsInfo gin.HandlerFunc = func(c *gin.Context) {
	station := c.Query("station")
	goodName := c.Query("good")
	if station == "" && goodName == "" {
		panic(errors.BadRequest("at least provide an argument: station or good"))
	}

	var goods []model.Good

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
	existGoodsMap := make(map[string]model.GoodsMapper)
	for _, mapper := range existGoodsMapper {
		existGoodsMap[mapper.ID] = mapper
	}

	var fullGoods []model.FullGood
	value := rediskit.GetValueFromRedis(strkit.Concat(hashkit.MD5(station), "_goods_list"), "")
	if value != "" {
		goodsIdList := mathutil.MustString(value)
		for _, goodId := range strings.Split(goodsIdList, ",") {
			var mapper model.GoodsMapper
			if goodMapper, ok := existGoodsMap[goodId]; ok {
				// filter known goods
				if show == "unknown" {
					continue
				}
				mapper = goodMapper
			} else {
				mapper = model.GoodsMapper{
					ID:      goodId,
					Name:    "Unknown",
					Station: station,
				}
			}
			fullGood := model.FullGood{
				Id:   goodId,
				Good: *mapperToGood(&mapper),
			}
			fullGoods = append(fullGoods, fullGood)
		}
	}

	c.JSON(http.StatusOK, fullGoods)
}

func mapperToGood(mapper *model.GoodsMapper) *model.Good {
	goodId := mapper.ID
	updateTimestamp := rediskit.GetValueFromRedis(strkit.Concat(goodId, "_update_time"), time.Now().Unix()+1000000)
	updateTimestampInt := mathutil.MustInt64(updateTimestamp)
	good := &model.Good{
		Name:            mapper.Name,
		Station:         mapper.Station,
		Price:           rediskit.GetIntFromRedis(strkit.Concat(goodId, "_price"), -1),
		NextTrend:       rediskit.GetIntFromRedis(strkit.Concat(goodId, "_next_trend"), 0),
		UpdateTime:      time.Unix(updateTimestampInt, 0).Format(time.DateTime),
		UpdateTimestamp: updateTimestampInt,
	}
	return good
}
