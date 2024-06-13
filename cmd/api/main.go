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

	e.GET("/demo/script.js", handler.Script)

	e.GET("/demo/style.css", handler.Stylesheet)

	e.GET("/v/:uuid", handler.Configuration)
	e.GET("/v/new", handler.Configuration)
	e.POST("/v/new", handler.IntroVideoCode)



	e.POST("/config", handler.GenerateCode)

	e.File("/", "internal/template/demo.html")

	e.Static("/", "public")

	e.Logger.Fatal(e.Start(":8080"))
}
