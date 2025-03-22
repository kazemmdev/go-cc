package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	ConnectToDB()

	r := chi.NewRouter()
	r.Route("/todos", func(r chi.Router) {
		r.Get("/", GetTodosHandler)
		r.Post("/", CreateTodoHandler)
		r.Get("/{id}", GetTodoHandler)
		r.Put("/{id}", UpdateTodoHandler)
		r.Delete("/{id}", DeleteTodoHandler)
	})

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
