package handler

import (
	"log"
	"net/http"
	"vizinhanca/internal/auth"
	"vizinhanca/internal/model"
	"vizinhanca/internal/repository"

	"github.com/gin-gonic/gin"
)

type CreateAlertInput struct {
	Title       string      `json:"title" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Category    string      `json:"category" binding:"required"`
	Location    model.Point `json:"location" binding:"required"`
}

func AlertHandler(c *gin.Context) {
	var input CreateAlertInput

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	claims, exists := c.Get("claims")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Claims not found in context"})
		return
	}

	userClaims, ok := claims.(*auth.AppClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims format"})
		return
	}

	alertToSave := model.Alert{
		Title:       input.Title,
		Description: input.Description,
		Category:    input.Category,
		Location:    input.Location,
		UserID:      userClaims.UserID,
	}

	if err := repository.CreateAlert(c.Request.Context(), &alertToSave); err != nil {
		log.Printf("ERROR: Failed to save alert: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create alert"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Alert created successfully",
		"alert":   alertToSave,
	})
}
