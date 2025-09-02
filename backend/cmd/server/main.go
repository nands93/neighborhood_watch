package main

import (
	"fmt"
	"vizinhanca/internal/database"
	"vizinhanca/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.ConnectToDB(); err != nil {
		fmt.Printf("Fatal error: failed to connect to database: %v", err)
	}

	// --- Rotas Abertas ---
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		api.GET("/health", handler.HealthCheck)
		api.POST("/auth/register", handler.RegisterUser)
		api.POST("/auth/login", handler.LoginUser)
	}

	// --- Rotas Protegidas ---
	authorized := router.Group("/api/v1")
	authorized.Use(handler.AuthMiddleware())
	{
		authorized.GET("/me", handler.GetCurrentUserProfile)
		authorized.POST("/alerts", handler.AlertHandler)
		authorized.GET("/alerts", handler.GetAlertsHandler)
	}

	router.Run()
}
