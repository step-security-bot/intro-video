package handler

import (
	"github.com/crocoder-dev/intro-video/internal"
	"github.com/crocoder-dev/intro-video/internal/config"
	"github.com/labstack/echo/v4"
)

func Stylesheet(c echo.Context) error {
	stylesheet := internal.Stylesheet{}

	stylesheetProps := internal.ProcessableFileProps{
		URL:    "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ForBiggerBlazes.mp4",
		Cta:    config.Cta{Enabled: true, TextContent: "Test", Type: config.DefaultCta},
		Bubble: config.Bubble{Enabled: true, TextContent: "Test", Type: config.DefaultBubble},
	}

	style, err := stylesheet.Process(stylesheetProps, internal.ProcessableFileOpts{ Preview: false })
	if err != nil {
		return err
	}

	return c.Blob(200, "text/css; charset=utf-8", []byte(style))
}
