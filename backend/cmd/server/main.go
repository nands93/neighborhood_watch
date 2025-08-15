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
	router := gin.Default()
	router.GET("/api/v1/health", handler.HealthCheck)
	router.POST("/api/v1/users", handler.RegisterUser)
	//router.POST("/api/v1/auth/token", handler.RegisterUser)
	router.Run()
}
