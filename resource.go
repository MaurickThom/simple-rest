package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// get handler for all request
func get(c echo.Context) error {
	cs := Contacts{}
	ct := c.Request().Header.Get("Accept")

	filter := c.QueryParam("city")
	if filter == "" {
		cs = getAll()
	} else {
		cs = getByCity(filter)
	}

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
	if co == nil {
		return c.NoContent(http.StatusNoContent)
	}

	ct := c.Request().Header.Get("Accept")

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

// create handler
func create(c echo.Context) error {
	co := Contact{}

	err := c.Bind(&co)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Estructura no válida: %v", err))
	}

	add(&co)

	ct := c.Request().Header.Get("Accept")

	switch ct {
	case "":
		fallthrough
	case "text/plain":
		return c.String(http.StatusCreated, co.String())
	case "application/json":
		return c.JSON(http.StatusCreated, co)
	case "text/xml":
		return c.XML(http.StatusCreated, co)
	default:
		return c.String(http.StatusBadRequest, "Debe especificar un content-type válido")
	}
}

// del handler
func del(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "El id debe ser numérico entero")
	}

	delete(id)
	return c.NoContent(http.StatusNoContent)
}

// upd handler
func upd(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "El id debe ser numérico entero")
	}

	co := Contact{}

	err = c.Bind(&co)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Estructura no válida: %v", err))
	}

	update(id, &co)
	ct := c.Request().Header.Get("Accept")

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
