package usecase

import (
	"fmt"
	"products_api/model"
	"products_api/repository"
)

type GetProductsUseCase struct {
	repo repository.ProductRepository
}

func NewGetProductsUseCase(r repository.ProductRepository) GetProductsUseCase {
	return GetProductsUseCase{
		repo: r,
	}
}

func (gp *GetProductsUseCase) Execute() ([]model.Product, error) {
	products := []model.Product{}

	products, err := gp.repo.ListAll()

	if err != nil {
		fmt.Printf("Unable to get products")
		return products, err
	}

	return products, nil

}
