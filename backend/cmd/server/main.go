package main

import (
	"fmt"
	"time"
	"vizinhanca/internal/database"
	"vizinhanca/internal/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.ConnectToDB(); err != nil {
		fmt.Printf("Fatal error: failed to connect to database: %v", err)
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		// Em produção, você deve restringir isso para o domínio exato do seu frontend.
		// Ex: AllowOrigins: []string{"https://www.meusite.com"},
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// --- Rotas Abertas ---

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
