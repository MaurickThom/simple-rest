package main

import (
	"github.com/labstack/echo"
	"net/http"
)

// get handler for all request
func get(c echo.Context) error {
	cs := getAll()
	return c.String(http.StatusOK, cs.String())
}
