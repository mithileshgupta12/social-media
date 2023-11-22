package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mithileshgupta12/social-media/handlers"
)

type Api struct{}

func (a *Api) Init(e *echo.Echo) {
	api := e.Group("/api")

	// Routes
	postRoutes(api)
}

func postRoutes(api *echo.Group) {
	postHandler := &handlers.PostHandler{}

	postGroup := api.Group("/posts")
	postGroup.GET("", postHandler.Get)
}
