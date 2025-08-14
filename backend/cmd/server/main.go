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
	router.Run()
}
