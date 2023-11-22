package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type PostHandler struct{}

func (p *PostHandler) Get(c echo.Context) error {
	type Response struct {
		Message string
	}

	return c.JSON(http.StatusOK, Response{
		Message: "Hey",
	})
}
