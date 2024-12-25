package usecase

import (
	"fmt"
	"products_api/model"
	"products_api/repository"
)

type GetProductsUseCase struct {
	repository repository.RepositoryInterface
}

func NewGetProductsUseCase(r repository.RepositoryInterface) GetProductsUseCase {
	return GetProductsUseCase{
		repository: r,
	}
}

func (gp *GetProductsUseCase) Execute() ([]model.Product, error) {
	products, err := gp.repository.ListAll()

	if err != nil {
		fmt.Printf("Unable to get products")
		return []model.Product{}, err
	}

	return products, nil

}
