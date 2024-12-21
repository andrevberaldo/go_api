package usecase

import "products_api/model"

type RepositoryInterface interface {
	Save(product model.Product) (string, error)
	ListAll() ([]model.Product, error)
	ListById(id int) (model.Product, error)
	Delete(id int) error
}
