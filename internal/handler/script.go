package handler

import (
	"github.com/crocoder-dev/intro-video/internal/script"
	"github.com/labstack/echo/v4"
)

func Script(c echo.Context) error {

	scriptProps := script.ScriptProps{
		Video:  script.Video{URL: "https://youtube.com/shorts/9K2ioP7aZcA?si=VcNgKCTEgyw7NipJ"},
		Cta:    script.Cta{Enabled: true, TextContent: "Test"},
		Bubble: script.Bubble{Enabled: true, TextContent: "Test"},
	}

	script, err := script.Process(scriptProps)
	if err != nil {
		return err
	}

	return c.Blob(200, "application/javascript; charset=utf-8", []byte(script))

}
