package routers

import "github.com/kmou424/resonance-dataserver/server/handlers"

func registerFetchRouters() {
	fetch := apiGroup.Group("/fetch")
	fetch.GET("/goods_info", handlers.GetGoodsInfo)
}
