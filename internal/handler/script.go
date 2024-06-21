package handler

import (
	"fmt"
	"strconv"

	"github.com/crocoder-dev/intro-video/internal"
	"github.com/crocoder-dev/intro-video/internal/config"
	"github.com/crocoder-dev/intro-video/internal/template"
	"github.com/labstack/echo/v4"
)

func Script(c echo.Context) error {
	script := internal.Script{}

	var scriptProps internal.ProcessableFileProps

	url := c.QueryParam(template.URL)

	bubbleEnabledRaw := c.QueryParam(template.BUBBLE_ENABLED)

	if bubbleEnabledRaw == "" {
		fmt.Println("bubbleEnabledRaw is empty")
	}

	var err error
	var bubbleEnabled bool

	if bubbleEnabledRaw != "" {
		bubbleEnabled, err = strconv.ParseBool(bubbleEnabledRaw)
		if err != nil {
			return err
		}
	}

	var bubbleTextContent string
	var bubbleType config.BubbleType

	if bubbleEnabled {
		bubbleTextContent = c.QueryParam(template.BUBBLE_TEXT)
		bubbleType, err = config.NewBubbleType(c.QueryParam(template.BUBBLE_TYPE))
		if err != nil {
			return err
		}
	}

	ctaEnabledRaw := c.QueryParam(template.CTA_ENABLED)

	var ctaEnabled bool

	if ctaEnabledRaw != "" {
		ctaEnabled, err = strconv.ParseBool(ctaEnabledRaw)
		if err != nil {
			return err
		}
	}

	var ctaTextContent string
	var ctaButtonType config.CtaButtonType

	if ctaEnabled {
		ctaTextContent = c.QueryParam(template.CTA_TEXT)
		ctaButtonType, err = config.NewCtaButtonType(c.QueryParam(template.CTA_TYPE))
		if err != nil {
			return err
		}
	}

	if url != "" {
		scriptProps = internal.ProcessableFileProps{
			URL: url,
			Bubble: config.Bubble{
				Enabled:     bubbleEnabled,
				TextContent: bubbleTextContent,
				Type:        bubbleType,
			},
			Cta: config.Cta{
				Enabled:     ctaEnabled,
				TextContent: ctaTextContent,
				Type:        ctaButtonType,
			},
		}
	} else {
		scriptProps = internal.ProcessableFileProps{
			URL:    "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ForBiggerBlazes.mp4",
			Cta:    config.Cta{Enabled: true, TextContent: "Test", Type: config.DefaultCta},
			Bubble: config.Bubble{Enabled: true, TextContent: "Test", Type: config.DefaultBubble},
		}
	}

	fmt.Println(scriptProps)

	s, err := script.Process(scriptProps, true)
	if err != nil {
		return err
	}

	return c.Blob(200, "application/javascript; charset=utf-8", []byte(s))
}
