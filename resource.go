package main

import (
	"github.com/labstack/echo"
	"net/http"
)

// get handler for all request
func get(c echo.Context) error {
	cs := getAll()
	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "":
		fallthrough
	case "text/plain":
		return c.String(http.StatusOK, cs.String())
	case "application/json":
		return c.JSON(http.StatusOK, cs)
	case "text/xml":
		return c.XML(http.StatusOK, cs)
	default:
		return c.String(http.StatusBadRequest, "Debe especificar un content-type válido")
	}
}
