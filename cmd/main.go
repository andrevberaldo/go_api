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
	productsRepo := repository.NewProductRepository(dbConnection)

	productController := controller.NewProductController(
		usecase.NewCreateProductUseCase(productsRepo),
		usecase.NewGetProductsUseCase(productsRepo),
		usecase.NewGetProductByIdUseCase(productsRepo),
	)

	server.GET("/health", HealthController.CheckHealth)
	server.POST("/products", middleware.AuthenticateJWT(), productController.CreateProduct)
	server.GET("/products", middleware.AuthenticateJWT(), productController.GetProducts)
	server.GET("/products/:id", middleware.AuthenticateJWT(), productController.GetProductById)

	server.Run(":3001")
}
