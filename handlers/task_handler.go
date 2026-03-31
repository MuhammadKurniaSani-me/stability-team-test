package handlers

import (
	"strconv"

	"stability-test-task-api/models"
	"stability-test-task-api/store"

	"github.com/gofiber/fiber/v2"
)

// GetTasks handles the HTTP GET request to retrieve all tasks.
// It fetches all tasks from the store and returns them as a JSON array.
func GetTasks(c *fiber.Ctx) error {
	tasks := store.GetAllTasks()
	return c.JSON(tasks)
}

// GetTask handles the HTTP GET request to retrieve a single task by its ID.
// It expects an 'id' parameter in the URL path.
func GetTask(c *fiber.Ctx) error {
	// Extract the 'id' parameter from the URL.
	idParam := c.Params("id")

	// Convert the ID parameter from a string to an integer.
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// If conversion fails, it means the ID is not a valid number.
		// Return a 400 Bad Request response with an error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid task id format",
		})
	}

	// Retrieve the task from the store using its ID.
	task := store.GetTaskByID(id)

	// If the task is not found (store returns nil), return a 404 Not Found response.
	if task == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "task not found",
		})
	}

	// If the task is found, return it as a JSON object with a 200 OK status.
	return c.JSON(task)
}

// CreateTask handles the HTTP POST request to create a new task.
// It expects a JSON body that can be unmarshalled into a Task struct.
func CreateTask(c *fiber.Ctx) error {
	// Create a new Task variable to hold the parsed request body.
	var task models.Task

	// Parse the JSON request body into the task struct.
	if err := c.BodyParser(&task); err != nil {
		// If parsing fails, return a 400 Bad Request response.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse request body",
		})
	}

	// Basic input validation: ensure the task has a title.
	if task.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "task title is required",
		})
	}

	// Add the new task to the store. The store will assign a unique ID.
	store.AddTask(&task)

	// Return the newly created task (including its new ID) as a JSON object
	// with a 201 Created status code.
	return c.Status(fiber.StatusCreated).JSON(task)
}

// DeleteTask handles the HTTP DELETE request to remove a task by its ID.
// It expects an 'id' parameter in the URL path.
func DeleteTask(c *fiber.Ctx) error {
	// Extract the 'id' parameter from the URL.
	idParam := c.Params("id")

	// Convert the ID parameter from a string to an integer.
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// If conversion fails, return a 400 Bad Request response.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid task id format",
		})
	}

	// Call the store to delete the task with the given ID.
	// This operation is idempotent; it doesn't fail if the ID doesn't exist.
	store.DeleteTask(id)

	// Return a success message with a 200 OK status.
	return c.JSON(fiber.Map{
		"message": "deleted",
	})
}
