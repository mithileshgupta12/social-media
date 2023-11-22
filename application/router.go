package application

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mithileshgupta12/social-media/routes"
)

type Router struct{}

func (r *Router) Init() *echo.Echo {
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	api := &routes.Api{}
	api.Init(e)

	return e
}
