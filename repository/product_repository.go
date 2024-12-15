package repository

import (
	"database/sql"
	"fmt"
	"products_api/model"
)

type ProductRepository struct {
	Connection *sql.DB
}

func NewProductRepository(dbconnection *sql.DB) ProductRepository {
	return ProductRepository{
		Connection: dbconnection,
	}
}

func (pr *ProductRepository) SaveProduct(product model.Product) (string, error) {
	query, err := pr.Connection.Prepare("INSERT INTO products(name, price) VALUES ($1, $2) RETURNING id")

	if err != nil {
		fmt.Printf("Error trying to save Product %s", err.Error())
		return "", err
	}

	var id int

	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		fmt.Printf("Error trying to save Product %s", err.Error())
		return "", err
	}

	query.Close()

	return fmt.Sprintf("/products/%d", id), nil
}
