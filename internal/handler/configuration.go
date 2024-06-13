package handler

import (
	"context"
	"fmt"
	"net/http"

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
	component := template.IntroVideoCode()
	return component.Render(context.Background(), c.Response().Writer)
}

func GenerateCode(c echo.Context) error {
	url := c.FormValue("url")
	bubbleEnabled := c.FormValue("bubble-enabled")
	bubbleTextContent := c.FormValue("bubble-text")
	bubbleType := c.FormValue("bubble-type")
	ctaEnabled := c.FormValue("cta-enabled")
	ctaTextContent := c.FormValue("cta-text")
	ctaType := c.FormValue("cta-tzpe")
	if url == "" {
		return c.String(http.StatusBadRequest, "Invalid data")
	}
	c.Set("url", url)
	c.Set("bubbleEnabled", bubbleEnabled)
	c.Set("bubbleTextContent", bubbleTextContent)
	c.Set("bubbleType", bubbleType)
	c.Set("ctaEnabled", ctaEnabled)
	c.Set("ctaTextContent", ctaTextContent)
	c.Set("ctaType", ctaType)

	// err, js := Script(c)
	// if err != nil {
	// 	return err
	// }

	// err, css := Stylesheet(c)
	// if err != nil {
	// 	return err
	// }

	// components := template.CodePreview(css, js)
	// return components.Render(context.Background(), c.Response().Writer)
	return nil
}
