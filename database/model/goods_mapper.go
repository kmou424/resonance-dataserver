package model

import "github.com/kmou424/resonance-dataserver/pojo/common"

const GoodsMapperTable = "goods_mappers"

type GoodsMapper struct {
	common.GoodBase

	ID string `json:"id" gorm:"primarykey"`
}
