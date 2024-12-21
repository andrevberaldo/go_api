package usecase

import (
	"fmt"
)

type DeleteProductUseCase struct {
	Repository RepositoryInterface
}

func NewDeleteProductUseCase(r RepositoryInterface) DeleteProductUseCase {
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
