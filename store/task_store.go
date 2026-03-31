package store

import (
	"stability-test-task-api/models"
	"sync"
)

var Tasks = []models.Task{
	{ID: 1, Title: "Learn Go", Done: false},
	{ID: 2, Title: "Build API", Done: false},
}

var (
	mu     sync.Mutex
	nextID = 3
)

func GetAllTasks() []models.Task {
	mu.Lock()
	defer mu.Unlock()
	return Tasks
}

func GetTaskByID(id int) *models.Task {
	mu.Lock()
	defer mu.Unlock()
	for _, t := range Tasks {
		if t.ID == id {
			task := t
			return &task
		}
	}
	return nil
}

func AddTask(task *models.Task) {
	mu.Lock()
	defer mu.Unlock()
	task.ID = nextID
	nextID++
	Tasks = append(Tasks, *task)
}

func DeleteTask(id int) {
	mu.Lock()
	defer mu.Unlock()
	for i, t := range Tasks {
		if t.ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			break
		}
	}
}
