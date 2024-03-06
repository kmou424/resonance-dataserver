package database

import (
	"encoding/json"
	"github.com/kmou424/resonance-dataserver/mapperdata"
	"github.com/kmou424/resonance-dataserver/model"
	"log"
)

func importData() {
	var goodMappers []model.GoodsMapper
	err := json.Unmarshal(mapperdata.ReadGoodsCitiesMapper(), &goodMappers)
	if err != nil {
		log.Fatal("can't parse mapper data", "error", err)
	}

	tx := Conn.Table(model.GoodsMapperTable).Create(goodMappers)
	if int(tx.RowsAffected) != len(goodMappers) {
		log.Fatal("import data failed, please check /database/data/goods_cities_mapper.json file")
	}
}
