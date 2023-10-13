package main

import (
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

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
