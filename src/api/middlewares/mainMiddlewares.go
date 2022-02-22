package middlewares

import (
	"github.com/labstack/echo/v4"
)

func SetMainMiddlewares(e *echo.Echo) {
	e.Use(ServerHeader)

}
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Mum-Server")
		return next(c)
	}
}
