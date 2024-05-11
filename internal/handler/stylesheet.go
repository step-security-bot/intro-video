package handler

import (
	"github.com/crocoder-dev/intro-video/internal/stylesheet"
	"github.com/labstack/echo/v4"
)


func Stylesheet(c echo.Context) error {

	stylesheetProps := stylesheet.StylesheetProps{
		Dimensions: stylesheet.Dimensions{
			Lwidth: 360,
			Lheight: 640,
			Swidth: 180,
			Sheight: 320,
		},
		Cta: stylesheet.Cta{},
		Bubble: stylesheet.Bubble{},
	}

	style, err := stylesheet.Process(stylesheetProps)
	if err != nil {
		return err
	}

	return c.Blob(200, "text/css; charset=utf-8", []byte(style))
}
