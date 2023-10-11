package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// InitializeDatabase initializes the database connection
func InitializeDatabase() (*sqlx.DB, error) {
	connStr := "user=postgres dbname=todo password=postgres sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
