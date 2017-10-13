package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func startRoutes() {
	e := echo.New()
	e.GET("/", index)
	e.GET("/contacts", get)
	e.GET("/contacts/:id", getID)

	e.Start(":8080")
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hola mundo")
}
