package pojo

import "github.com/kmou424/resonance-dataserver/pojo/common"

type Good struct {
	common.GoodBase
	GoodExtra
}

type GoodExtra struct {
	Price           int    `json:"price"`
	NextTrend       int    `json:"next_trend"`
	UpdateTime      string `json:"update_time"`
	UpdateTimestamp int64  `json:"update_timestamp"`
}

type FullGood struct {
	ID string `json:"id"`
	Good
}
