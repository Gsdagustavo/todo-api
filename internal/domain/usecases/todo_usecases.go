package usecases

import (
	"todo-api/internal/domain/entities"
	"todo-api/internal/domain/errors"
	"todo-api/internal/domain/repositories"
)

type TodoUseCase struct {
	r repositories.TodoRepository
}

func NewTodoUseCase(r repositories.TodoRepository) *TodoUseCase {
	return &TodoUseCase{r}
}

func (u *TodoUseCase) AddTodo(request entities.AddTodoRequest) (*entities.AddTodoResponse, error) {
	if request.Title == "" {
		return nil, errors.ErrInvalidTitle
	}

	return u.r.AddTodo(request)
}
