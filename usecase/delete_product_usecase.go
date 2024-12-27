package usecase

import (
	"fmt"
)

type IDeleteteRepository interface {
	Delete(id int) error
}

type DeleteProductUseCase struct {
	Repository IDeleteteRepository
}

func NewDeleteProductUseCase(r IDeleteteRepository) DeleteProductUseCase {
	return DeleteProductUseCase{
		Repository: r,
	}
}

func (u *DeleteProductUseCase) Execute(id int) error {
	err := u.Repository.Delete(id)

	if err != nil {
		fmt.Printf("Unable to delete product id = %v", id)
		return err
	}

	return nil
}
