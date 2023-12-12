package repositories

import "golang-todo-app/models"

var todos = []models.Todo{{Id: 1, Description: "Fake DB", Completed: false}}

func GetTodos() []models.Todo {
	return todos
}

func StoreTodo(todo models.Todo) (createdTodo models.Todo, err error) {
	todos = append(todos, todo)
	return todo, nil
}
