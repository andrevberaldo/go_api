package repository

import (
	"database/sql"
	"fmt"
	"products_api/model"
)

var (
	SAVE_PRODUCT_SQL      = "INSERT INTO products(name, price) VALUES ($1, $2) RETURNING id"
	LIST_ALL_PRODUCTS_SQL = "SELECT id, name, price FROM products"
	LIST_BY_ID_SQL        = "SELECT id, name, price FROM products WHERE id=$1"
	DELETE_BY_ID_SQL      = "DELETE FROM products WHERE id=$1"
)

type ProductRepository struct {
	Connection *sql.DB
}

func NewProductRepository(dbconnection *sql.DB) ProductRepository {
	return ProductRepository{
		Connection: dbconnection,
	}
}

func (r *ProductRepository) Save(product model.Product) (int, error) {
	var id int

	err := r.Connection.
		QueryRow(SAVE_PRODUCT_SQL, product.Name, product.Price).
		Scan(&id)

	if err != nil {
		fmt.Printf("Error trying to save Product %s", err.Error())
		return -1, err
	}

	return id, nil
}

func (r *ProductRepository) ListAll() ([]model.Product, error) {
	products := []model.Product{}

	rows, err := r.Connection.Query(LIST_ALL_PRODUCTS_SQL)

	if err != nil {
		fmt.Print("Unable to load products from database")
		return products, err
	}

	defer rows.Close()

	var product model.Product

	for rows.Next() {
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
		)

		if err != nil {
			fmt.Printf("Cannot parse row into a product entity %v\n", err)
			return products, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepository) ListById(id int) (model.Product, error) {
	var product model.Product

	err := r.Connection.QueryRow(LIST_BY_ID_SQL, id).Scan(&product.ID, &product.Name, &product.Price)

	if err != nil {
		fmt.Printf("Unable to load product_id=%v from database due to %v", id, err.Error())
		return model.Product{}, err
	}

	return product, nil
}

func (r *ProductRepository) Delete(id int) error {
	_, err := r.Connection.Exec(DELETE_BY_ID_SQL, id)

	if err != nil {
		fmt.Printf("Unable to delete product id %v", id)
		return err
	}

	return nil
}
