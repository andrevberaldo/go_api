package usecase

import (
	"fmt"
	"products_api/model"
)

type GetProductByIdUseCase struct {
	repository RepositoryInterface
}

func NewGetProductByIdUseCase(r RepositoryInterface) GetProductByIdUseCase {
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
