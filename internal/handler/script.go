package handler

import (
	"github.com/crocoder-dev/intro-video/internal"
	"github.com/crocoder-dev/intro-video/internal/config"
	"github.com/labstack/echo/v4"
)

func Script(c echo.Context) error {
	script := internal.Script{}

	scriptProps := internal.ProcessableFileProps{
		URL:    "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ForBiggerBlazes.mp4",
		Theme:  config.DefaultTheme,
		Cta:    config.Cta{Enabled: true, TextContent: "Test"},
		Bubble: config.Bubble{Enabled: true, TextContent: "Test"},
	}

	s, err := script.Process(scriptProps, internal.ProcessableFileOpts{Preview: false})
	if err != nil {
		return err
	}

	return c.Blob(200, "application/javascript; charset=utf-8", []byte(s))
}
