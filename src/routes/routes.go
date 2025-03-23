package routes

import (
	"github.com/go-chi/chi/v5"
	"go-cc/src/handlers"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/todos", func(r chi.Router) {
		r.Get("/", handlers.GetTodosHandler)
		r.Post("/", handlers.CreateTodoHandler)
		r.Get("/{id}", handlers.GetTodoHandler)
		r.Put("/{id}", handlers.UpdateTodoHandler)
		r.Delete("/{id}", handlers.DeleteTodoHandler)
	})

	return r
}
