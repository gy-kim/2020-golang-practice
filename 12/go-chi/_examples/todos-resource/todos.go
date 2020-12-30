package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

type todosResource struct{}

// Routes creates a REST router for the todos resource
func (rs todosResource) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	r.Get("/", rs.List)
	r.Post("/", rs.Create)
	r.Put("/", rs.Delete)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", rs.Get)
		r.Put("/", rs.Update)
		r.Delete("/", rs.Delete)
		r.Get("/sync", rs.Sync)
	})

	return r
}

func (rs todosResource) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todos list of stuff.."))
}

func (rs todosResource) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todos create"))
}

func (rs todosResource) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todo get"))
}

func (rs todosResource) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todo update"))
}

func (rs todosResource) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todo delete"))
}

func (rs todosResource) Sync(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todo sync"))
}
