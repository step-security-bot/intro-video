package main

import (
	"github.com/crocoder-dev/intro-video/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	e.GET("/script.js", handler.Script)

	e.GET("/style.css", handler.Stylesheet)

	e.Logger.Fatal(e.Start(":8080"))
}
