package main

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
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

// getID handler
func getID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "El id debe ser numérico entero")
	}

	co := getByID(id)
	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "":
		fallthrough
	case "text/plain":
		return c.String(http.StatusOK, co.String())
	case "application/json":
		return c.JSON(http.StatusOK, co)
	case "text/xml":
		return c.XML(http.StatusOK, co)
	default:
		return c.String(http.StatusBadRequest, "Debe especificar un content-type válido")
	}
}
