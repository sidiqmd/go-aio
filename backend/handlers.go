package main

import (
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `db:"id"`
	Title     string `db:"title"`
	Completed bool   `db:"completed"`
}

// GetAllTodosHandler handles the GET /api/todos route
func getAllTodosHandler(c *fiber.Ctx) error {
	todos, err := getAllTodo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(todos)
}

// GetAllTodo queries all Todo items from the database
func getAllTodo() ([]Todo, error) {
	todos := []Todo{}
	err := db.Select(&todos, "SELECT id, title, completed FROM todos")
	if err != nil {
		return nil, err
	}
	return todos, nil
}

// GetTodoById handles the GET /api/todos/:id route
func getTodoByIdHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	todo, err := getTodoById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(todo)
}

// GetTodoById queries a single Todo item from the database
func getTodoById(id string) (Todo, error) {
	var todo Todo
	err := db.Get(&todo, "SELECT id, title, completed FROM todos WHERE id=$1", id)
	if err != nil {
		return Todo{}, err
	}
	return todo, nil
}

// CreateTodoHandler handles the POST /api/todos route
func createTodoHandler(c *fiber.Ctx) error {
	type request struct {
		Title string `json:"title"`
	}
	// New struct
	r := new(request)
	// Parse body into struct
	if err := c.BodyParser(r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Create a new Todo item
	todo, err := createTodoItem(r.Title)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Return created Todo item as JSON
	return c.JSON(todo)
}

// CreateTodoItem creates a new Todo item in the database
func createTodoItem(title string) (Todo, error) {
	var todo Todo
	err := db.Get(&todo, "INSERT INTO todos (title, completed) VALUES ($1, false) RETURNING id, title, completed", title)
	if err != nil {
		return Todo{}, err
	}
	return todo, nil
}

// UpdateTodoHandler handles the PUT /api/todos/:id route
func updateTodoHandler(c *fiber.Ctx) error {
	type request struct {
		Title string `json:"title"`
	}
	// New struct
	r := new(request)
	// Parse body into struct
	if err := c.BodyParser(r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	id := c.Params("id")
	// Create a new Todo item
	todo, err := updateTodoItem(id, r.Title)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Return created Todo item as JSON
	return c.JSON(todo)
}

// UpdateTodoItem updates a Todo item in the database
func updateTodoItem(id string, title string) (Todo, error) {
	var todo Todo
	err := db.Get(&todo, "UPDATE todos SET title=$1 WHERE id=$2 RETURNING id, title, completed", title, id)
	if err != nil {
		return Todo{}, err
	}
	return todo, nil
}

// DeleteTodoHandler handles the DELETE /api/todos/:id route
func toggleTodoCompletedHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	todo, err := getTodoById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	todo, err = toggleTodoCompleted(id, todo.Completed)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(todo)
}

// ToggleTodoComplete toggles the completed status of a Todo item in the database
func toggleTodoCompleted(id string, completed bool) (Todo, error) {
	var todo Todo
	err := db.Get(&todo, "UPDATE todos SET completed = $1 WHERE id=$2 RETURNING id, title, completed", !completed, id)
	if err != nil {
		return Todo{}, err
	}
	return todo, nil
}

// DeleteTodoHandler handles the DELETE /api/todos/:id route
func deleteTodoHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	err := deleteTodoItem(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusOK)
}

// DeleteTodoItem deletes a Todo item from the database
func deleteTodoItem(id string) error {
	_, err := db.Exec("DELETE FROM todos WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
