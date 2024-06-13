package handler

import (
	"github.com/crocoder-dev/intro-video/internal"
	"github.com/crocoder-dev/intro-video/internal/config"
	"github.com/labstack/echo/v4"
)

func Script(c echo.Context) (string, error) {
	url := c.Get("url").(string)

	script := internal.Script{}

	scriptProps := internal.ProcessableFileProps{
		URL:    url,
		Cta:    config.Cta{Enabled: true, TextContent: "Test", Type: config.DefaultCta},
		Bubble: config.Bubble{Enabled: true, TextContent: "Test", Type: config.DefaultBubble},
	}

	s, err := script.Process(scriptProps)
	if err != nil {
		return "", err
	}

	return s, nil
}
