package repositories

import "github.com/kmou424/resonance-dataserver/database"

func InitRepositories() {
	AuthKey = &AuthKeyRepository{database.Conn}
	GoodsMapper = &GoodsMapperRepository{database.Conn}
}
