package main

import (
	"products_api/controller"
	"products_api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	HealthController := controller.NewHealthController()
	ProductController := controller.NewProductController()

	server.GET("/health", HealthController.CheckHealth)
	server.POST("/products", middleware.AuthenticateJWT(), ProductController.CreateProduct)

	server.Run(":3001")
}
