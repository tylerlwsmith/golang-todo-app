package repositories

import "golang-todo-app/models"

var todos = []models.Todo{}

func GetTodos() []models.Todo {
	return todos
}

func StoreTodo(todo models.Todo) (createdTodo models.Todo, err error) {
	todos = append(todos, todo)
	return todo, nil
}

func DeleteTodo(id int) {
	filtered := []models.Todo{}
	for _, todo := range todos {
		if todo.Id != id {
			filtered = append(filtered, todo)
		}
	}

	todos = filtered
}
