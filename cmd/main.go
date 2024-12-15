package main

import (
	"products_api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	HealthController := controller.NewHealthController()
	server.GET("/health", HealthController.CheckHealth)

	server.Run(":3001")
}
