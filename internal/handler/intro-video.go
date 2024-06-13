package handler

import (
	"context"
	"net/http"

	"github.com/crocoder-dev/intro-video/internal/template"
	"github.com/labstack/echo/v4"
)

type IntroVideoData struct {
	Js  string `json:"js"`
	Css string `json:"css"`
}

func IntroVideo(c echo.Context) error {
	components := template.IntroVideoPage()
	return components.Render(context.Background(), c.Response().Writer)
}

func GenerateCode(c echo.Context) error {
	url := c.FormValue("url")
	if url == "" {
		return c.String(http.StatusBadRequest, "Invalid data")
	}
	c.Set("url", url)

	js, err := Script(c)
	if err != nil {
		return err
	}

	css, err := Stylesheet(c)
	if err != nil {
		return err
	}

	components := template.CodeTextareas(css, js)
	return components.Render(context.Background(), c.Response().Writer)
}
