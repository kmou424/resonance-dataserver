package handlers

import "github.com/gin-gonic/gin"

var Health gin.HandlerFunc = func(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}
