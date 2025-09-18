package router

import (
	"fmt"
	"log"
	"net/http"
	"todo-api/internal/http/handlers"

	"github.com/gorilla/mux"
)

type MuxRouter struct {
	router *mux.Router
}

func NewMuxRouter() *MuxRouter {
	return &MuxRouter{router: mux.NewRouter()}
}

func (r *MuxRouter) ServeHTTP(port string) {
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r.router)
	if err != nil {
		panic(err)
	}
}

func (r *MuxRouter) StartServer(handlers handlers.Handlers, port string) {
	log.Printf("Server is listening on port %s", port)
	r.registerHandlers(handlers)
	r.ServeHTTP(port)
}

func (r *MuxRouter) registerHandlers(h handlers.Handlers) {
	r.router.HandleFunc("/todos/add", handlers.Handle(h.TodoHandler.AddTodo))
}
