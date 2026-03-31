package models

// Task represents the data model for a single task in the system.
type Task struct {
	// ID is the unique identifier for the task.
	// The `json:"id"` struct tag specifies how this field should be
	// named when the struct is serialized to or deserialized from JSON.
	ID int `json:"id"`

	// Title is a short description of the task.
	// The `json:"title"` struct tag maps this field to "title" in JSON.
	Title string `json:"title"`

	// Done indicates whether the task has been completed.
	// The `json:"done"` struct tag maps this field to "done" in JSON.
	Done bool `json:"done"`
}
