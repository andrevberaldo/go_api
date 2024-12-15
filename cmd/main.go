package main

import (
	"products_api/controller"
	"products_api/db"
	"products_api/middleware"
	"products_api/repository"
	"products_api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	HealthController := controller.NewHealthController()
	productController := controller.NewProductController(usecase.NewProductUseCase(repository.NewProductRepository(dbConnection)))

	server.GET("/health", HealthController.CheckHealth)
	server.POST("/products", middleware.AuthenticateJWT(), productController.CreateProduct)

	server.Run(":3001")
}
