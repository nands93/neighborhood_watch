package handler

import (
	"log"
	"net/http"
	"vizinhanca/internal/auth"
	"vizinhanca/internal/model"
	"vizinhanca/internal/repository"

	"github.com/gin-gonic/gin"
)

type RegisterUserInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func RegisterUser(c *gin.Context) {
	var input RegisterUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	validEmail := auth.ValidateEmail(input.Email)
	if !validEmail {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email format",
		})
		return
	}

	strengthRules := auth.DefaultPasswordStrength()
	passwordErrors := auth.ValidatePasswordStrength(input.Password, strengthRules)
	if len(passwordErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Password validation failed",
			"details": passwordErrors,
		})
		return
	}

	userToSave := model.User{
		Username: input.Username,
		Email:    input.Email,
	}

	hashedPassword, err := auth.GenerateArgon2Hash(input.Password)
	if err != nil {
		log.Printf("ERROR: Failed to hash password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	userToSave.Password = hashedPassword
	if err := repository.CreateUser(c.Request.Context(), &userToSave); err != nil {
		log.Printf("ERROR: Failed to register user in repository: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"username": userToSave.Username,
		"email":    userToSave.Email,
		"status":   "user registered successfully",
	})
}

func GetCurrentUserProfile(c *gin.Context) {
	// Pegamos os claims que o nosso middleware adicionou ao contexto.
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Claims not found in context"})
		return
	}

	// Retornamos os claims como JSON para o cliente.
	c.JSON(http.StatusOK, gin.H{"user_profile": claims})
}
