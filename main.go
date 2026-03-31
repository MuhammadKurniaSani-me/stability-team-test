package main

import (
	"stability-test-task-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize a new Fiber application.
	app := fiber.New()

	// Define the routes for the Task API.
	// Each route is mapped to a handler function from the 'handlers' package.

	// GET /tasks - Retrieves a list of all tasks.
	app.Get("/tasks", handlers.GetTasks)
	// GET /tasks/:id - Retrieves a single task by its ID.
	app.Get("/tasks/:id", handlers.GetTask)
	// POST /tasks - Creates a new task.
	app.Post("/tasks", handlers.CreateTask)
	// DELETE /tasks/:id - Deletes a task by its ID.
	app.Delete("/tasks/:id", handlers.DeleteTask)

	// Start the server and listen for incoming requests on port 3000.
	// This is a blocking call that keeps the application running.
	app.Listen(":3000")
}
