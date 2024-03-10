package model

const GoodsMapperTable = "goods_mappers"

type GoodsMapper struct {
	ID      string `json:"id" gorm:"primarykey"`
	Name    string `json:"name"`
	Station string `json:"station"`
}
