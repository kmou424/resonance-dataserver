package repositories

import "github.com/kmou424/resonance-dataserver/database"

func InitRepositories() {
	GoodsMapper = &GoodsMapperRepository{database.Conn}
}
