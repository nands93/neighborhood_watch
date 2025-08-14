package main

import (
	"vizinhanca/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/api/v1/health", handler.HealthCheck)
	router.Run()
}
