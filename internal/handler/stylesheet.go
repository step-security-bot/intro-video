package handler

import (
	"github.com/crocoder-dev/intro-video/internal/stylesheet"
	"github.com/labstack/echo/v4"
)


func Stylesheet(c echo.Context) error {

	stylesheetProps := stylesheet.StylesheetProps{
		Cta: stylesheet.Cta{},
		Bubble: stylesheet.Bubble{},
	}

	style, err := stylesheet.Process(stylesheetProps)
	if err != nil {
		return err
	}

	return c.Blob(200, "text/css; charset=utf-8", []byte(style))
}
