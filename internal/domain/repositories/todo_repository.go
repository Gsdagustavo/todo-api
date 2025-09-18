package repositories

import "todo-api/internal/domain/entities"

type TodoRepository interface {
	AddTodo(request entities.AddTodoRequest) (*entities.AddTodoResponse, error)
}
