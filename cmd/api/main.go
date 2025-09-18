package main

import (
	"todo-api/internal/domain/usecases"
	"todo-api/internal/infrastructure/persistence"
	"todo-api/internal/infrastructure/repositories"
)

func main() {
	db, err := persistence.OpenDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	todoRepository := repositories.NewMySQLTodoRepository(db)
	_ = usecases.NewTodoUseCase(todoRepository)
}
