package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// routes generates our routes and attaches them to handlers, using the chi router
// note that we return type http.Handler, and not *chi.Mux; since chi.Mux satisfies
// the interface requirements for http.Handler, it makes sense to return the type
// that is part of the standard library.
func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)

	mux.Get("/users/login", app.Login)
	mux.Post("/users/login", app.Login)

	return mux
}
