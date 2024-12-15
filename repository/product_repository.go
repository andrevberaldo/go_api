package repository

import "products_api/model"

type ProductRepository struct {
}

func NewProductRepository() ProductRepository {
	return ProductRepository{}
}

func (pr *ProductRepository) SaveProduct(product model.Product) (model.Product, error) {
	created := model.Product{
		ID:    333,
		Name:  product.Name,
		Price: product.Price,
	}

	return created, nil
}
