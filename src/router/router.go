package router

import (
	"myapp/src/api"
	"myapp/src/api/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	adminGroup := e.Group("/admin")
	cookieGrop := e.Group("/cookie")
	jwtGroup := e.Group("/jwt")

	middlewares.SetMainMiddlewares(e)
	middlewares.SetAdminMiddlewares(adminGroup)
	middlewares.SetCookieMiddlewares(cookieGrop)
	middlewares.SetJwtMiddlewares(jwtGroup)

	api.MainGroup(e)
	//set Group route
	api.AdminGroup(adminGroup)
	api.CookieGroup(cookieGrop)
	api.JwtGroup(jwtGroup)

	return e
}
