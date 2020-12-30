package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

type usersResource struct{}

// Routes crease a REST router for the todos resource
func (rs usersResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.List)
	r.Post("/", rs.Create)
	r.Put("/", rs.Delete)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", rs.Get)
		r.Put("/", rs.Update)
		r.Delete("/", rs.Delete)
	})
	return r
}

func (rs usersResource) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aaa list of stuff.."))
}

func (rs usersResource) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aaa create"))
}

func (rs usersResource) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aaa get"))
}

func (rs usersResource) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aaa update"))
}

func (rs usersResource) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aaa delete"))
}
