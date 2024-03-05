package main

import (
	"github.com/kmou424/resonance-dataserver/cache"
	"github.com/kmou424/resonance-dataserver/database"
	"github.com/kmou424/resonance-dataserver/database/repositories"
	"github.com/kmou424/resonance-dataserver/server"
)

func init() {
	cache.InitRedis()
	database.InitSQLite()
	repositories.InitRepositories()
}

func main() {
	server.Run()
}
