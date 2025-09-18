package usecases

import (
	"todo-api/internal/domain/repositories"
)

type TodoUseCase struct {
	r repositories.TodoRepository
}

func NewTodoUseCase(r repositories.TodoRepository) *TodoUseCase {
	return &TodoUseCase{r}
}
