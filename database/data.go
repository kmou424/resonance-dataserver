package database

import (
	"encoding/json"
	"github.com/kmou424/resonance-dataserver/database/data"
	"github.com/kmou424/resonance-dataserver/database/model"
	"log"
)

func importData() {
	importMapperData()
	importAuthKeysData()
}

func importMapperData() {
	var goodMappers []model.GoodsMapper
	err := json.Unmarshal(data.ReadGoodsCitiesMapper(), &goodMappers)
	if err != nil {
		log.Fatal("can't parse mapper data", "error", err)
	}

	tx := Conn.Table(model.GoodsMapperTable).Create(goodMappers)
	if int(tx.RowsAffected) != len(goodMappers) {
		log.Fatal("import data failed, please check /database/data/goods_cities_mapper.json file")
	}
}

func importAuthKeysData() {
	var authKeys []model.AuthKey
	err := json.Unmarshal(data.ReadAuthKeys(), &authKeys)
	if err != nil {
		log.Fatal("can't parse mapper data", "error", err)
	}

	tx := Conn.Table(model.AuthKeyTable).Create(authKeys)
	if int(tx.RowsAffected) != len(authKeys) {
		log.Fatal("import data failed, please check /database/data/auth_keys.json file")
	}
}
