package handler

import (
	"context"
	"fmt"
	"strconv"

	"github.com/crocoder-dev/intro-video/internal"
	"github.com/crocoder-dev/intro-video/internal/config"
	"github.com/crocoder-dev/intro-video/internal/template"
	"github.com/labstack/echo/v4"
)

func Configuration(c echo.Context) error {

	uuid := c.Param("uuid")

	if uuid == "" {
		fmt.Println("uuid is empty")
	} else {
		fmt.Println("uuid is:", uuid)
	}

	bubleOptions := []template.BubbleOption{
		{Caption: "Default Bubble", Value: config.DefaultBubble, Selected: true},
		{Caption: "Custom Bubble", Value: config.CustomBubble, Selected: false},
	}

	ctaOptions := []template.CtaOption{
		{Caption: "Default CTA", Value: config.DefaultCta, Selected: true},
		{Caption: "Custom CTA", Value: config.CustomCta, Selected: false},
	}

	component := template.Configuration(bubleOptions, ctaOptions)
	return component.Render(context.Background(), c.Response().Writer)
}

func IntroVideoCode(c echo.Context) error {
	url := c.FormValue(template.URL)

	bubbleEnabledRaw := c.FormValue(template.BUBBLE_ENABLED)

	fmt.Println("bubbleEnabledRaw:", bubbleEnabledRaw)

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
		bubbleTextContent = c.FormValue(template.BUBBLE_TEXT)
		bubbleType, err = config.NewBubbleType(c.FormValue(template.BUBBLE_TYPE))
		if err != nil {
			return err
		}
	}

	ctaEnabledRaw := c.FormValue(template.CTA_ENABLED)

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
		ctaTextContent = c.FormValue(template.CTA_TEXT)
		ctaButtonType, err = config.NewCtaButtonType(c.FormValue(template.CTA_TYPE))
		if err != nil {
			return err
		}
	}

	processableFileProps := internal.ProcessableFileProps{
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

	js, err := internal.Script{}.Process(processableFileProps, false)
	if err != nil {
		return err
	}

	css, err := internal.Stylesheet{}.Process(processableFileProps)
	if err != nil {
		return err
	}

	component := template.IntroVideoCode(js, css)
	return component.Render(context.Background(), c.Response().Writer)
}
