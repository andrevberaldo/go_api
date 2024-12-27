package usecase

import (
	"fmt"
	"products_api/model"
)

type IRepositoryListAll interface {
	ListAll() ([]model.Product, error)
}

type GetProductsUseCase struct {
	repository IRepositoryListAll
}

func NewGetProductsUseCase(r IRepositoryListAll) GetProductsUseCase {
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
