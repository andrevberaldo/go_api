package usecase

import (
	"fmt"
	"products_api/model"
)

type GetProductsUseCase struct {
	repository RepositoryInterface
}

func NewGetProductsUseCase(r RepositoryInterface) GetProductsUseCase {
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
