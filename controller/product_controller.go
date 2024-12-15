package controller

import (
	"fmt"
	"net/http"
	"products_api/model"
	"products_api/usecase"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) ProductController {
	return ProductController{
		ProductUsecase: usecase,
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

	created, err := pc.ProductUsecase.CreateProduct(product)

	if err != nil {
		fmt.Printf("Failure on creating a new Product %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, created)
}
