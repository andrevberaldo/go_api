package controller

import (
	"fmt"
	"net/http"
	"products_api/model"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
}

func NewProductController() ProductController {
	return ProductController{}
}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	var product model.Product

	err := ctx.BindJSON(&product)

	if err != nil {
		fmt.Printf("Not possible to parse request JSON")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}

	ctx.JSON(http.StatusCreated, product)
}
