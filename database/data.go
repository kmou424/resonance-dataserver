package database

import (
	"encoding/json"
	"github.com/kmou424/resonance-dataserver/mapperdata"
	"github.com/kmou424/resonance-dataserver/model"
	"log"
)

func importData() {
	content, err := mapperdata.Data.ReadFile("files/goods_cities_mapper.json")
	if err != nil {
		log.Fatal("can't read mapper data", "error", err)
	}
	var goodMappers []model.GoodsMapper
	err = json.Unmarshal(content, &goodMappers)
	if err != nil {
		log.Fatal("can't parse mapper data", "error", err)
	}

	tx := Conn.Table(model.GoodsMapperTable).Create(goodMappers)
	if int(tx.RowsAffected) != len(goodMappers) {
		log.Fatal("import data failed, please check /database/data/goods_cities_mapper.json file")
	}
}
