package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// initRoutes initialises the routes for the server
func initRoutes(e *echo.Echo) {
	e.GET("/", homeHandler)
	e.GET("/results", resultsHandler)
}

// homeHandler handlers route to /
func homeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

// resultsHandler handles routes to /results
func resultsHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "results", nil)
}
