package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kmou424/resonance-dataserver/server/engine"
	"github.com/kmou424/resonance-dataserver/server/handlers"
)

var apiGroup *gin.RouterGroup

func Register() {
	apiGroup = engine.E.Group("/api")
	registerFetchRouters()
	registerAdminRouters()
}

func registerBaseRouters() {
	apiGroup.GET("/health", handlers.Health)
}
