package usecase

import (
	"fmt"
	"products_api/model"
)

type IRepositoryListById interface {
	ListById(id int) (model.Product, error)
}

type GetProductByIdUseCase struct {
	repository IRepositoryListById
}

func NewGetProductByIdUseCase(r IRepositoryListById) GetProductByIdUseCase {
	return GetProductByIdUseCase{
		repository: r,
	}
}

func (u *GetProductByIdUseCase) Execute(id int) (model.Product, error) {
	product, err := u.repository.ListById(id)

	if err != nil {
		fmt.Printf("Unable to get products")
		return model.Product{}, err
	}

	return product, nil

}
