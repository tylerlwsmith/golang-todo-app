package repositories

import "golang-todo-app/models"

var tasks = []models.Task{}

func GetTasks() []models.Task {
	return tasks
}

func StoreTask(task models.Task) (createdTask models.Task, err error) {
	tasks = append(tasks, task)
	return task, nil
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
