package handler

import (
	"net/http"
	"vizinhanca/internal/auth"
	"vizinhanca/internal/repository"

	"github.com/gin-gonic/gin"
)

// DTO for login input
type LoginUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginUser(c *gin.Context) {
	var input LoginUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := repository.GetUserAuth(c.Request.Context(), input.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user credentials"})
		return
	}

	checkHash, err := auth.CheckPasswordHash(input.Password, user.Password)
	if err != nil || !checkHash {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password credentials"})
		return
	}

	token, err := auth.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token + "\n",
	})
}
