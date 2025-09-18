package repositories

import (
	"database/sql"
	"todo-api/internal/domain/entities"
)

type MySQLTodoRepository struct {
	database *sql.DB
}

func NewMySQLTodoRepository(database *sql.DB) *MySQLTodoRepository {
	return &MySQLTodoRepository{database}
}

func (r *MySQLTodoRepository) AddTodo(request entities.AddTodoRequest) (*entities.AddTodoResponse, error) {
	const query = "INSERT INTO todos (title, description, is_completed) VALUES (?, ?, ?)"

	result, err := r.database.Exec(query, request.Title, request.Description, request.Completed)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &entities.AddTodoResponse{Id: id}, nil
}
