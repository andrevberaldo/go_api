package usecase

import (
	"fmt"
	"products_api/repository"
)

type DeleteProductUseCase struct {
	repo repository.ProductRepository
}

func NewDeleteProductUseCase(r repository.ProductRepository) DeleteProductUseCase {
	return DeleteProductUseCase{
		repo: r,
	}
}

func (pu *DeleteProductUseCase) Execute(id int) error {
	err := pu.repo.DeleteProduct(id)

	if err != nil {
		fmt.Printf("Unable to delete product id = %v", id)
		return err
	}

	return nil
}
