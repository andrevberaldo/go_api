package usecase

import (
	"fmt"
	"products_api/repository"
)

type DeleteProductUseCase struct {
	Repository repository.RepositoryInterface
}

func NewDeleteProductUseCase(r repository.RepositoryInterface) DeleteProductUseCase {
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
