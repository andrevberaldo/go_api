package usecase

import (
	"fmt"
	"products_api/model"
	"products_api/repository"
	"strconv"
)

type CreateProductUsecase struct {
	Repository repository.RepositoryInterface
}

func NewCreateProductUseCase(repository repository.RepositoryInterface) CreateProductUsecase {
	return CreateProductUsecase{
		Repository: repository,
	}
}

func (u *CreateProductUsecase) Execute(product model.Product) (string, error) {
	id, err := u.Repository.Save(product)

	if err != nil {
		fmt.Printf("Error creating new Product")
		return "", err
	}

	return "/api/products/" + strconv.Itoa(id), nil
}
