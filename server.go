package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	username := c.QueryParam("username")
	name := c.QueryParam("name")
	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Your username is %s \n Your name %s", username, name))
	}
	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"username": username,
			"name":     name,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"err": "invalid data type",
	})
}
func halo(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "Admin page")
}

//Custom Middleware
// ServerHeader middleware adds a `Server` header to the response.
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Mum-Server")
		return next(c)
	}
}

func main() {
	e := echo.New()
	e.Use(ServerHeader)
	g := e.Group("/admin")

	//**** middleware ******

	//e.Use(middleware.Logger())
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	//Auth
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if username == "mum" && password == "1234" {
			return true, nil
		}
		return false, nil
	}))

	//**** Route ******
	g.GET("/main", mainAdmin)
	e.GET("/users/:data", getUser)
	e.GET("/", halo)

	e.Start(":1323")
}
