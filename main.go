package main

import (
	"github.com/DaveSaah/risky/views"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("./static"))
	e.Renderer = views.CreateTemplates("./views")
	e.Logger.SetLevel(log.DEBUG)

	initRoutes(e)

	// start server with logger
	e.Logger.Fatal(e.Start(":8000"))
}
