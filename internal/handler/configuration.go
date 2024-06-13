package handler

import (
	"context"
	"fmt"

	"github.com/crocoder-dev/intro-video/internal/template"
	"github.com/labstack/echo/v4"
)

func Configuration(c echo.Context) error {

	uuid := c.Param("uuid")

	if uuid == "" {
		fmt.Println("uuid is empty")
	} else {
		fmt.Println("uuid is:", uuid)
	}

	component := template.Configuration()
	return component.Render(context.Background(), c.Response().Writer)
}

func IntroVideoCode(c echo.Context) error {
	component := template.IntroVideoCode()
	return component.Render(context.Background(), c.Response().Writer)
}
