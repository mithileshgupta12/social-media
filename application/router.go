package application

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mithileshgupta12/social-media/handlers"
)

type Router struct{}

func (router *Router) Init() *chi.Mux {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)

	// Routes
	r.Route("/", userRoutes)
	r.Route("/posts", postRoutes)

	return r
}

func userRoutes(r chi.Router) {
	userHandler := &handlers.UserHandler{}

	r.Get("/login", userHandler.Login)
}

func postRoutes(r chi.Router) {
	postHandler := &handlers.PostHandler{}

	r.Get("/", postHandler.Get)
}
