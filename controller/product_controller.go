package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"products_api/model"
	"products_api/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	createProductUseCase  usecase.CreateProductUsecase
	getProductsUseCase    usecase.GetProductsUseCase
	getProductByIdUseCase usecase.GetProductByIdUseCase
	deleteProductUseCase  usecase.DeleteProductUseCase
}

func NewProductController(
	createUsecase usecase.CreateProductUsecase,
	getProductsUseCase usecase.GetProductsUseCase,
	getProductByIdUseCase usecase.GetProductByIdUseCase,
	deleteProductUseCase usecase.DeleteProductUseCase,
) ProductController {

	return ProductController{
		createProductUseCase:  createUsecase,
		getProductsUseCase:    getProductsUseCase,
		getProductByIdUseCase: getProductByIdUseCase,
		deleteProductUseCase:  deleteProductUseCase,
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
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		fmt.Printf("An product id must be provided")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "An product id must be provided",
		})
		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		fmt.Printf("Could not parse string to int")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse string to int",
		})
		return
	}

	product, err := pc.getProductByIdUseCase.Execute(productId)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Failure on get Product %v", err.Error())
			ctx.JSON(http.StatusNotFound, gin.H{})
			return
		}

		fmt.Printf("Failure on get Product %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (pc *ProductController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		fmt.Printf("An product id must be provided")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "An product id must be provided",
		})
		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		fmt.Printf("Could not parse string to int")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse string to int",
		})
		return
	}

	err = pc.deleteProductUseCase.Execute(productId)

	if err != nil {
		fmt.Printf("Could not delete product")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete product",
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{})

}
