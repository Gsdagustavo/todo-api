package usecases

import (
	"strings"
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
	validatedTitle := strings.TrimSpace(request.Title)
	validatedDescription := strings.TrimSpace(request.Description)

	if validatedTitle == "" {
		return nil, errors.ErrInvalidTitle
	}

	req := entities.AddTodoRequest{
		Title:       request.Title,
		Description: validatedDescription,
	}

	return u.r.AddTodo(req)
}
