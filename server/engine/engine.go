package engine

import "github.com/gin-gonic/gin"

var E *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	E = gin.Default()
}
