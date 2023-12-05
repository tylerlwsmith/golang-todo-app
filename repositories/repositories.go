package repositories

import "golang-todo-app/models"

var todos []models.Todo

func init() {
	todos = append(todos, models.Todo{Id: 1, Description: "Fake DB", Done: false})
}

func GetTodos() []models.Todo {
	return todos
}
