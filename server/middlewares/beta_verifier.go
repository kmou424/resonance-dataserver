package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/kmou424/resonance-dataserver/internal/kits/hashkit"
	"github.com/kmou424/resonance-dataserver/server/errors"
)

const betaToken = "30d7c4d89aca2a893c3efec08f2b2f46"

func GetBetaVerifier() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			c.Abort()
			panic(errors.BadRequest("you must provide a token if you want to try beta apis"))
		}
		if hashkit.MD5(token) != betaToken {
			c.Abort()
			panic(errors.BadRequest("your token is invalid"))
		}

		c.Next()
	}
}
