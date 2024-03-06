package model

type Good struct {
	Name    string `json:"name"`
	Station string `json:"station"`

	Price           int    `json:"price"`
	NextTrend       int    `json:"next_trend"`
	UpdateTime      string `json:"update_time"`
	UpdateTimestamp int64  `json:"update_timestamp"`
}

type FullGood struct {
	Id string `json:"id"`
	Good
}
