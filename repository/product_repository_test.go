package repository_test

import (
	"errors"
	"products_api/model"
	"products_api/repository"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	// given
	givenProduct := model.Product{Name: "Test Product", Price: 100.0}
	givenProductIdDB := 1

	// Create DB Mock using the sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error on creating DB mock: %v", err)
	}
	defer db.Close()

	// Crease repository using the mock DB
	repo := repository.NewProductRepository(db)

	t.Run("should call the DB and return the resource URL", func(t *testing.T) {
		// Stubing the DB return depending on the args
		sql := regexp.QuoteMeta(repository.SAVE_PRODUCT_SQL)
		mock.ExpectQuery(sql).
			WithArgs(givenProduct.Name, givenProduct.Price).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(givenProductIdDB))

		// when
		url, err := repo.Save(givenProduct)

		// then
		assert.NoError(t, err)
		assert.Equal(t, givenProductIdDB, url)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Expectations not achieved: %v", err)
		}
	})

	t.Run("should return the error when connection goes wrong and an empty string", func(t *testing.T) {
		// Crease repository using the mock DB
		errorMessage := "fake connection error"

		// Stubing the DB return depending on the args
		sql := regexp.QuoteMeta(repository.SAVE_PRODUCT_SQL)
		mock.ExpectQuery(sql).
			WithArgs(givenProduct.Name, givenProduct.Price).
			WillReturnError(errors.New(errorMessage))

		// when
		url, err := repo.Save(givenProduct)

		// then
		assert.Error(t, err)
		assert.Equal(t, -1, url)
		assert.Equal(t, errorMessage, err.Error())

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Expectations not achieved: %v", err)
		}
	})
}

func TestListAll(t *testing.T) {
	// given
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error on creating DB mock: %v", err)
	}
	defer db.Close()

	repo := repository.NewProductRepository(db)

	t.Run("should return all the products", func(t *testing.T) {
		// given
		sql := regexp.QuoteMeta(repository.LIST_ALL_PRODUCTS_SQL)
		mock.ExpectQuery(sql).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price"}).
				AddRow(1, "Product 1", 100.0).
				AddRow(2, "Product 2", 200.0))

		// when
		products, err := repo.ListAll()

		// then
		assert.NoError(t, err)
		assert.Len(t, products, 2)
		assert.Equal(t, "Product 1", products[0].Name)
		assert.Equal(t, float32(100.0), products[0].Price)
		assert.Equal(t, "Product 2", products[1].Name)
		assert.Equal(t, float32(200.0), products[1].Price)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Expectations not achieved: %v", err)
		}
	})

	t.Run("should return an empty products list and the error if db connection fails", func(t *testing.T) {
		// given
		dbConnectionErrorMessage := "connection failed"

		sql := regexp.QuoteMeta(repository.LIST_ALL_PRODUCTS_SQL)
		mock.ExpectQuery(sql).
			WillReturnError(errors.New(dbConnectionErrorMessage))

		// when
		products, err := repo.ListAll()

		// then
		assert.Error(t, err)
		assert.Len(t, products, 0)
		assert.Equal(t, dbConnectionErrorMessage, err.Error())

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Expectations not achieved: %v", err)
		}
	})

	t.Run("should return an empty products list and the error if parse to product entity fails", func(t *testing.T) {
		// given
		repo := repository.NewProductRepository(db)

		sql := regexp.QuoteMeta(repository.LIST_ALL_PRODUCTS_SQL)
		mock.ExpectQuery(sql).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price"}).
				AddRow(nil, "Product 1", 100.0))

		// when
		products, err := repo.ListAll()

		// then
		assert.Error(t, err)
		assert.Len(t, products, 0)
		assert.Equal(t, "sql: Scan error on column index 0, name \"id\": converting NULL to int is unsupported", err.Error())

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Expectations not achieved: %v", err)
		}
	})

}

func TestListById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error on creating DB mock: %v", err)
	}
	defer db.Close()

	repo := repository.NewProductRepository(db)

	t.Run("should return the product without error", func(t *testing.T) {
		// given
		sql := regexp.QuoteMeta(repository.LIST_BY_ID_SQL)
		mock.ExpectQuery(sql).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price"}).
				AddRow(1, "Product 1", 100.0))

		// when
		product, err := repo.ListById(1)

		// then
		assert.NoError(t, err)
		assert.Equal(t, 1, product.ID)
		assert.Equal(t, "Product 1", product.Name)
		assert.Equal(t, float32(100.0), product.Price)

		// Verifique se o mock foi chamado corretamente
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Expectations not achieved: %v", err)
		}
	})

	t.Run("should return an empty product and the error if connection fails", func(t *testing.T) {
		// given
		fakeErrorMessage := "something went wrong"
		sql := regexp.QuoteMeta(repository.LIST_BY_ID_SQL)
		mock.ExpectQuery(sql).
			WillReturnError(errors.New(fakeErrorMessage))

		// when
		product, err := repo.ListById(1)

		// then
		assert.Error(t, err)
		assert.Equal(t, fakeErrorMessage, err.Error())
		assert.Equal(t, 0, product.ID)
		assert.Equal(t, "", product.Name)
		assert.Equal(t, float32(0), product.Price)

		// Verifique se o mock foi chamado corretamente
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Expectations not achieved: %v", err)
		}
	})

}

func TestDelete(t *testing.T) {
	// given
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error on creating DB mock: %v", err)
	}
	defer db.Close()

	repo := repository.NewProductRepository(db)

	t.Run("should delete the product and return no error", func(t *testing.T) {
		// given
		sql := regexp.QuoteMeta(repository.DELETE_BY_ID_SQL)
		mock.ExpectExec(sql).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// when
		err = repo.Delete(1)

		// then
		assert.NoError(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Expectations not achieved: %v", err)
		}
	})

	t.Run("should return an error if something went wrong", func(t *testing.T) {
		// given
		fakeErrorMessage := "something went wrong"
		sql := regexp.QuoteMeta(repository.DELETE_BY_ID_SQL)
		mock.ExpectExec(sql).
			WillReturnError(errors.New(fakeErrorMessage))

		// when
		err = repo.Delete(1)

		// then
		assert.Error(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Expectations not achieved: %v", err)
		}
	})

}
