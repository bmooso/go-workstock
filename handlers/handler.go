package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type Handler struct {
}

// InitRoutes Initializes the routes for the handler.
func (h Handler) InitRoutes(e *echo.Echo) {
	internal := e.Group("/hello")
	internal.GET("", h.hello)
}

func (h Handler) hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World!")
}
