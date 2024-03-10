package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/kmou424/resonance-dataserver/server/errors"
)

func getErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpError, ok := err.(errors.HttpError)
				if ok {
					c.JSON(httpError.StatusCode, map[string]any{
						"error": httpError.Error.Error(),
					})
				} else {
					panic(err)
				}
			}
		}()

		c.Next()
	}
}
