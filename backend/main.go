package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

// Define a global database connection
var db *sqlx.DB

func main() {
	// Initialize the database
	var err error
	db, err = InitializeDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// db = db

	app := fiber.New()

	app.Get("/todos", getAllTodosHandler)
	app.Get("/todos/:id", getTodoByIdHandler)
	app.Post("/todos", createTodoHandler)
	app.Put("/todos/:id", updateTodoHandler)
	app.Patch("/todos/:id", toggleTodoCompletedHandler)
	app.Delete("/todos/:id", deleteTodoHandler)

	// Start the Fiber server
	port := 8080
	fmt.Printf("Server is running on :%d...\n", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
