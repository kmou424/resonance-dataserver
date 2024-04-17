package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kmou424/resonance-dataserver/database/repositories"
	"github.com/kmou424/resonance-dataserver/server/errors"
	"strings"
)

func getAuthVerifier() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasSuffix(c.FullPath(), "/api/health") {
			c.Next()
			return
		}
		uuid := c.Query("uuid")
		verified := repositories.AuthKey.HasUUID(uuid)
		if !verified {
			c.Abort()
			panic(errors.Unauthorized(fmt.Sprintf("provided uuid is invalid")))
		}

		c.Next()
	}
}

func getAdminVerifier() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasSuffix(c.FullPath(), "/api/admin") {
			if uuid := c.Query("uuid"); !repositories.AuthKey.IsAdmin(uuid) {
				c.Abort()
				panic(errors.Forbidden(fmt.Sprintf("permission denied")))
			}
		}

		c.Next()
	}
}
