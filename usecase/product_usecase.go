package usecase

import (
	"fmt"
	"products_api/model"
	"products_api/repository"
)

type ProductUsecase struct {
	Repository repository.ProductRepository
}

func NewProductUseCase(repository repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		Repository: repository,
	}
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	created, err := pu.Repository.SaveProduct(product)

	if err != nil {
		fmt.Printf("Error creating new Product")
		return model.Product{
			ID:    0,
			Name:  "",
			Price: 0,
		}, err
	}

	return created, nil
}
