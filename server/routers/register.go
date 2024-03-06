package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kmou424/resonance-dataserver/server/engine"
	"github.com/kmou424/resonance-dataserver/server/middlewares"
)

var apiGroup *gin.RouterGroup
var betaApiGroup *gin.RouterGroup

func Register() {
	apiGroup = engine.E.Group("/api")
	betaApiGroup = apiGroup.Group("/beta", middlewares.GetBetaVerifier())
	registerFetchRouters()
}
