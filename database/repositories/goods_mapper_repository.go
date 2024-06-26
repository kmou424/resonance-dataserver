package repositories

import (
	"github.com/kmou424/resonance-dataserver/database/model"
	"gorm.io/gorm"
)

var GoodsMapper *GoodsMapperRepository

type GoodsMapperRepository struct {
	db *gorm.DB
}

func (repo *GoodsMapperRepository) Find(name string, station string) (goodsMapper []model.GoodsMapper) {
	tx := repo.db.Table(model.GoodsMapperTable)

	if name != "" {
		tx = tx.Where("name = ?", name)
	}
	if station != "" {
		tx = tx.Where("station = ?", station)
	}

	tx.Find(&goodsMapper)

	return
}
