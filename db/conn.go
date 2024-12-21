package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DB_CONNECTION_INFO"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Connected to " + os.Getenv("DB_NAME"))

	return db, nil
}
