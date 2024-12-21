package usecase

import (
	"fmt"
	"products_api/model"
)

type CreateProductUsecase struct {
	Repository RepositoryInterface
}

func NewCreateProductUseCase(repository RepositoryInterface) CreateProductUsecase {
	return CreateProductUsecase{
		Repository: repository,
	}
}

func (u *CreateProductUsecase) Execute(product model.Product) (string, error) {
	location, err := u.Repository.Save(product)

	if err != nil {
		fmt.Printf("Error creating new Product")
		return "", err
	}

	return location, nil
}
