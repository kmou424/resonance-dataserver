package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kmou424/resonance-dataserver/database/repositories"
	"github.com/kmou424/resonance-dataserver/server/errors"
)

func getAuthVerifier() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Query("uuid")
		verified := repositories.AuthKey.Has(uuid)
		if !verified {
			c.Abort()
			panic(errors.Unauthorized(fmt.Sprintf("provided uuid is invalid")))
		}

		c.Next()
	}
}
