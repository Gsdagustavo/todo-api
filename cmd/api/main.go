package main

import (
	"os"
	"todo-api/internal/domain/usecases"
	"todo-api/internal/http/handlers"
	"todo-api/internal/http/handlers/todo"
	"todo-api/internal/http/router"
	"todo-api/internal/infrastructure/persistence"
	"todo-api/internal/infrastructure/repositories"
)

var serverPort = os.Getenv("SERVER_PORT")

func main() {
	db, err := persistence.OpenDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	/// REPOSITORY
	todoRepository := repositories.NewMySQLTodoRepository(db)

	/// USE CASES
	todoUseCases := usecases.NewTodoUseCase(todoRepository)

	/// HANDLERS
	todoHandler := todo.NewTodoHandler(todoUseCases)

	handlers := handlers.Handlers{TodoHandler: todoHandler}

	router := router.NewMuxRouter()
	router.StartServer(handlers, serverPort)
}
