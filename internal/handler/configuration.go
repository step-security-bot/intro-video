package handler

import (
	"context"
	"fmt"

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
	fmt.Println(
		"url", c.FormValue(template.URL), "\n",
		"bubbleEnabled", c.FormValue(template.BUBBLE_ENABLED), "\n",
		"bubbleText", c.FormValue(template.BUBBLE_TEXT), "\n",
		"bubbleType", c.FormValue(template.BUBBLE_TYPE), "\n",
		"ctaEnabled", c.FormValue(template.CTA_ENABLED), "\n",
		"ctaText", c.FormValue(template.CTA_TEXT), "\n",
		"ctaType", c.FormValue(template.CTA_TYPE),
	)



	url := c.FormValue(template.URL)

	bubbleEnabledRaw := c.FormValue(template.BUBBLE_ENABLED)

	var bubbleEnabled bool
	var err error

	if bubbleEnabledRaw == "" {
		bubbleEnabled = false
	} else if bubbleEnabledRaw == "true" {
		bubbleEnabled = true
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

	var ctaEnabled bool

	ctaEnabledRaw := c.FormValue(template.CTA_ENABLED)

	if ctaEnabledRaw == "" {
		ctaEnabled = false
	} else if ctaEnabledRaw == "true" {
		ctaEnabled = true
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

	previewScript, err := internal.Script{}.Process(processableFileProps, internal.ProcessableFileOpts{ Preview: true})
	previewScript = "<script>" + previewScript + "</script>"

	previewStyle, err := internal.Stylesheet{}.Process(processableFileProps, internal.ProcessableFileOpts{ Preview: true })
	previewStyle = "<style>" + previewStyle + "</style>"

	js, err := internal.Script{}.Process(processableFileProps, internal.ProcessableFileOpts{})
	if err != nil {
		return err
	}

	css, err := internal.Stylesheet{}.Process(processableFileProps, internal.ProcessableFileOpts{})
	if err != nil {
		return err
	}

	component := template.IntroVideoPreview(js, css, previewScript, previewStyle)
	return component.Render(context.Background(), c.Response().Writer)
}
