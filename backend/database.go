package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// InitializeDatabase initializes the database connection
func InitializeDatabase() (*sqlx.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	sslMode := os.Getenv("DB_SSLMODE")
	connStr := "host=" + host + " port=" + port + " user=" + user + " dbname=" + dbname + " password=" + password + " sslmode=" + sslMode
	// connStr := "host=localhost port=5432 user=postgres dbname=todo password=postgres sslmode=disable"

	fmt.Printf("connStr: %s\n", connStr)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
