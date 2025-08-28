package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"vizinhanca/internal/auth"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			return
		}
		claims, err := auth.ValidateJWT(token)
		if err != nil {
			log.Printf("Token validation error: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		var userID int64

		if claims.UserID != 0 {
			userID = claims.UserID
		} else if claims.Subject != "" {
			if id, err := strconv.ParseInt(claims.Subject, 10, 64); err == nil {
				userID = id
			}
		}

		if userID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user id not found in token"})
		}

		c.Set("userID", userID)
		c.Set("claims", claims)
		c.Next()
	}
}
