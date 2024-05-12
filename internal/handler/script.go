package handler

import (
	"github.com/crocoder-dev/intro-video/internal/config"
	"github.com/crocoder-dev/intro-video/internal"
	"github.com/labstack/echo/v4"
)

func Script(c echo.Context) error {
	script := internal.Script{}


	scriptProps := internal.ProcessableFileProps{
		Video:  config.Video{URL: "https://youtube.com/shorts/9K2ioP7aZcA?si=VcNgKCTEgyw7NipJ"},
		Cta:    config.Cta{Enabled: true, TextContent: "Test", Type: config.DefaultCta},
		Bubble: config.Bubble{Enabled: true, TextContent: "Test", Type: config.DefaultBubble},
	}

	s, err := script.Process(scriptProps)
	if err != nil {
		return err
	}

	return c.Blob(200, "application/javascript; charset=utf-8", []byte(s))

}
