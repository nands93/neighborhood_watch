package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AlertHandler(c *gin.Context) {
	userID := c.GetInt64("userID")
	if userID == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
		return
	}

}
