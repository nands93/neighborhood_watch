package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AlertHandler(c *gin.Context) {
	v, ok := c.Get("claims")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no claims in context"})
		return
	}

	var userID string

}
