package handler

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/crocoder-dev/intro-video/internal/template"
)

func Configuration(c echo.Context) error {
	component := template.Configuration()

	return component.Render(context.Background(), c.Response().Writer)
}
