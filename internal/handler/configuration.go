package handler

import (
	"context"
	"net/http"

	"github.com/crocoder-dev/intro-video/internal/template"
	"github.com/labstack/echo/v4"
)

func Configuration(c echo.Context) error {
	component := template.Configuration()

	return component.Render(context.Background(), c.Response().Writer)
}

func GenerateCode(c echo.Context) error {
	url := c.FormValue("url")
	bubbleEnabled := c.FormValue("bubble-enabled")
	bubbleTextContent := c.FormValue("bubble-text")
	ctaEnabled := c.FormValue("cta-enabled")
	ctaTextContent := c.FormValue("cta-text")
	if url == "" {
		return c.String(http.StatusBadRequest, "Invalid data")
	}
	c.Set("url", url)
	c.Set("bubbleEnabled", bubbleEnabled)
	c.Set("bubbleTextContent", bubbleTextContent)
	c.Set("ctaEnabled", ctaEnabled)
	c.Set("ctaTextContent", ctaTextContent)

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
