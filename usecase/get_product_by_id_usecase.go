package usecase

import (
	"fmt"
	"products_api/model"
	"products_api/repository"
)

type GetProductByIdUseCase struct {
	repo repository.ProductRepository
}

func NewGetProductByIdUseCase(r repository.ProductRepository) GetProductByIdUseCase {
	return GetProductByIdUseCase{
		repo: r,
	}
}

func (gp *GetProductByIdUseCase) Execute(id int) (model.Product, error) {
	product, err := gp.repo.GetProductById(id)

	if err != nil {
		fmt.Printf("Unable to get products")
		return model.Product{}, err
	}

	return product, nil

}
