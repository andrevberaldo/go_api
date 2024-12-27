package usecase

import (
	"fmt"
	"products_api/model"
	"strconv"
)

type ICreateRepository interface {
	Save(product model.Product) (int, error)
}

type CreateProductUsecase struct {
	Repository ICreateRepository
}

func NewCreateProductUseCase(repository ICreateRepository) CreateProductUsecase {
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
