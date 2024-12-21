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

func (pr *ProductRepository) Save(product model.Product) (string, error) {
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

	return fmt.Sprintf("/api/products/%d", id), nil
}

func (pr *ProductRepository) ListAll() ([]model.Product, error) {
	products := []model.Product{}

	rows, err := pr.Connection.Query("SELECT id, name, price FROM products")

	if err != nil {
		fmt.Print("Unable to load products from database")
		return products, err
	}

	var product model.Product

	for rows.Next() {
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
		)

		if err != nil {
			fmt.Print("Unable to load products from database")
			return products, err
		}

		products = append(products, product)
	}

	rows.Close()

	return products, nil
}

func (pr *ProductRepository) ListById(id int) (model.Product, error) {
	query, err := pr.Connection.Prepare("SELECT id, name, price FROM products WHERE id=$1")

	if err != nil {
		fmt.Printf("Unable to load product_id=%v from database due to %v", id, err.Error())
		return model.Product{}, err
	}

	var product model.Product

	err = query.QueryRow(id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)

	if err != nil {
		fmt.Printf("Unable to load product_id=%v from database due to %v", id, err.Error())
		return model.Product{}, err
	}

	query.Close()

	return product, nil
}

func (pr *ProductRepository) Delete(id int) error {
	_, err := pr.Connection.Exec("DELETE FROM products WHERE id=$1", id)

	if err != nil {
		fmt.Printf("Unable to delete product id %v", id)
		return err
	}

	return nil
}
