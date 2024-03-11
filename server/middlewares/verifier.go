package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kmou424/resonance-dataserver/database/repositories"
	"github.com/kmou424/resonance-dataserver/server/errors"
)

func getAuthVerifier() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Query("show") == "unknown" {
			if c.FullPath() == "/api/fetch/full_goods_info" {
				c.Next()
				return
			}
		}

		uuid := c.Query("uuid")
		verified := repositories.AuthKey.Has(uuid)
		if !verified {
			c.Abort()
			panic(errors.Unauthorized(fmt.Sprintf("provided uuid is invalid")))
		}

		c.Next()
	}
}
