package usecase

import (
	"fmt"
	"products_api/model"
	"products_api/repository"
)

type CreateProductUsecase struct {
	Repository repository.ProductRepository
}

func NewCreateProductUseCase(repository repository.ProductRepository) CreateProductUsecase {
	return CreateProductUsecase{
		Repository: repository,
	}
}

func (pu *CreateProductUsecase) Execute(product model.Product) (string, error) {
	location, err := pu.Repository.SaveProduct(product)

	if err != nil {
		fmt.Printf("Error creating new Product")
		return "", err
	}

	return location, nil
}
