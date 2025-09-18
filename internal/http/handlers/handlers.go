package handlers

import (
	"net/http"
	"todo-api/internal/http/core"
	"todo-api/internal/http/handlers/todo"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) *core.ApiError

type Handlers struct {
	TodoHandler *todo.TodoHandler
}

func Handle(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if apiErr := h(w, r); apiErr != nil {
			apiErr.WriteHTTP(w)
		}
	}
}
