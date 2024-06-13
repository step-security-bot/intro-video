package handler

import (
	"github.com/crocoder-dev/intro-video/internal"
	"github.com/crocoder-dev/intro-video/internal/config"
	"github.com/labstack/echo/v4"
)

func Stylesheet(c echo.Context) (string, error) {
	url := c.Get("url").(string)

	stylesheet := internal.Stylesheet{}

	stylesheetProps := internal.ProcessableFileProps{
		URL:    url,
		Cta:    config.Cta{Enabled: true, TextContent: "Test"},
		Bubble: config.Bubble{Enabled: true, TextContent: "Test"},
	}

	style, err := stylesheet.Process(stylesheetProps)
	if err != nil {
		return "", err
	}

	return style, nil
}
