package main

import (
	"log"
	"net/http"

	"github.com/DaveSaah/risky/controllers"
	"github.com/labstack/echo/v4"
)

// initRoutes initialises the routes for the server
func initRoutes(e *echo.Echo) {
	e.GET("/", homeHandler)
	e.GET("/results", resultsHandler)
	e.POST("/simulate", simulateHandler)
}

// homeHandler handlers route to /
func homeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

// resultsHandler handles routes to /results
func resultsHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "results", nil)
}

// simulateHandler handles all logic from validation to simulation
func simulateHandler(c echo.Context) error {
	tradeInfo, err := controllers.Validate(c)
	if err != nil {
		return err
	}

	log.Printf("Trade info: %v\n", tradeInfo)

	return nil
}
