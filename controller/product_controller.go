package controller

import (
	"fmt"
	"net/http"
	"products_api/model"
	"products_api/usecase"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	createProductUseCase usecase.CreateProductUsecase
	getProductsUseCase   usecase.GetProductsUseCase
}

func NewProductController(createUsecase usecase.CreateProductUsecase, getProductsUseCase usecase.GetProductsUseCase) ProductController {
	return ProductController{
		createProductUseCase: createUsecase,
		getProductsUseCase:   getProductsUseCase,
	}
}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	var product model.Product

	err := ctx.ShouldBindJSON(&product)

	if err != nil {
		fmt.Printf("Failed to parse body to JSON %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	location, err := pc.createProductUseCase.Execute(product)

	if err != nil {
		fmt.Printf("Failure on creating a new Product %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"url": "http://localhost:3001" + location,
	})
}

func (pc *ProductController) GetProducts(ctx *gin.Context) {
	products, err := pc.getProductsUseCase.Execute()

	if err != nil {
		fmt.Printf("Failure on creating a new Product %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, products)
}
