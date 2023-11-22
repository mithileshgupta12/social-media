package application

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Route struct{}

func LoadRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/posts", postRoutes)

	return r
}

func postRoutes(r chi.Router) {
	//
}
