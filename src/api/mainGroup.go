package api

import (
	"myapp/src/api/handlers"

	"github.com/labstack/echo/v4"
)

func MainGroup(e *echo.Echo) {
	e.GET("/login", handlers.Login)
	e.GET("/users/:data", handlers.GetUser)
	e.GET("/", handlers.Halo)
}
