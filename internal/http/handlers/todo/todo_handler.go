package todo

import (
	"encoding/json"
	"errors"
	"net/http"
	"todo-api/internal/domain/entities"
	errors2 "todo-api/internal/domain/errors"
	"todo-api/internal/domain/usecases"
	"todo-api/internal/http/core"
)

type TodoHandler struct {
	TodoUseCases *usecases.TodoUseCase
}

func NewTodoHandler(useCase *usecases.TodoUseCase) *TodoHandler {
	return &TodoHandler{TodoUseCases: useCase}
}

func (h *TodoHandler) AddTodo(w http.ResponseWriter, r *http.Request) *core.ApiError {
	if err := ValidateRequestMethod(r, http.MethodPost); err != nil {
		return err.WithError("todo handler | add todo")
	}

	var request entities.AddTodoRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return (&core.ApiError{
			Code:    http.StatusBadRequest,
			Message: core.ErrBadRequest.Error(),
			Err:     core.ErrBadRequest,
		}).WithError("todo handler | add todo")
	}

	response, err := h.TodoUseCases.AddTodo(request)
	if err != nil {
		if errors.Is(err, core.ErrBadRequest) {
			return (&core.ApiError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
				Err:     err,
			}).WithError("todo handler | add todo")
		}

		if errors.Is(err, errors2.ErrInvalidTitle) {
			return (&core.ApiError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
				Err:     err,
			}).WithError("todo handler | add todo")
		}

		return (&core.ApiError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Err:     err,
		}).WithError("todo handler | add todo")
	}

	json.NewEncoder(w).Encode(response)

	return nil
}

func ValidateRequestMethod(r *http.Request, allowedMethod string) *core.ApiError {
	if r.Method != allowedMethod {
		return &core.ApiError{
			Code:    http.StatusMethodNotAllowed,
			Message: core.ErrMethodNotAllowed.Error(),
			Err:     core.ErrMethodNotAllowed,
		}
	}

	return nil
}
