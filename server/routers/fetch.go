package routers

import "github.com/kmou424/resonance-dataserver/server/handlers"

func registerFetchRouters() {
	fetch := apiGroup.Group("/fetch")
	fetch.GET("/goods_info", handlers.GetGoodsInfo)
	fetch.GET("/full_goods_info", handlers.GetFullGoodsInfo)

	//fetchBeta := betaApiGroup.Group("/fetch")
}
