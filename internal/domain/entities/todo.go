package entities

import "time"

type Todo struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
	Completed   bool      `json:"completed"`
}

type AddTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type AddTodoResponse struct {
	Id int64 `json:"id"`
}
