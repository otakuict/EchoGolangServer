package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
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
