package main

import (
	"time"
	"todo-api/internal/infrastructure/persistence"
)

func main() {
	db, err := persistence.OpenDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	time.Sleep(5 * time.Hour)
}
