package repositories

import (
	"errors"
	"golang-todo-app/models"
)

var tasks = []models.Task{}

func GetTasks() []models.Task {
	return tasks
}

func GetTask(id int) (task models.Task, err error) {
	found := false
	for _, t := range tasks {
		if t.Id == id {
			found = true
			task = t
			break
		}
	}

	if found != true {
		err = errors.New("No task found")
	}

	return task, err
}

func StoreTask(task models.Task) (createdTask models.Task, err error) {
	tasks = append(tasks, task)
	return task, nil
}

func UpdateTask(id int, task models.Task) (updatedTask models.Task, err error) {
	found := false
	for i, t := range tasks {
		if t.Id == id {
			found = true
			tasks[i] = task
			break
		}
	}

	if found != true {
		err = errors.New("No task found")
	}

	return task, err
}

func DeleteTask(id int) {
	filtered := []models.Task{}
	for _, task := range tasks {
		if task.Id != id {
			filtered = append(filtered, task)
		}
	}

	tasks = filtered
}
