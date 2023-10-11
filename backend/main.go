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

	baseUrl := "/api/todo"

	app.Get(baseUrl, getAllTodosHandler)
	app.Get(baseUrl+"/:id", getTodoByIdHandler)
	app.Post(baseUrl, createTodoHandler)
	app.Put(baseUrl+"/:id", updateTodoHandler)
	app.Patch(baseUrl+"/:id", toggleTodoCompletedHandler)
	app.Delete(baseUrl+"/:id", deleteTodoHandler)

	// Start the Fiber server
	port := 8080
	fmt.Printf("Server is running on :%d...\n", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
