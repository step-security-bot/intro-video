package handler

import (
	"github.com/crocoder-dev/intro-video/internal/config"
	"github.com/crocoder-dev/intro-video/internal"
	"github.com/labstack/echo/v4"
)


func Stylesheet(c echo.Context) error {

	stylesheet := internal.Stylesheet{}

	stylesheetProps := internal.ProcessableFileProps{
		Cta: config.Cta{ Enabled: true, TextContent: "Test" },
		Bubble: config.Bubble{ Enabled: true, TextContent: "Test" },
	}

	style, err := stylesheet.Process(stylesheetProps)
	if err != nil {
		return err
	}

	return c.Blob(200, "text/css; charset=utf-8", []byte(style))
}
