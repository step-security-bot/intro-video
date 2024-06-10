package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type IntroVideoData struct {
	Js  string `json:"js"`
	Css string `json:"css"`
}

func IntroVideo(c echo.Context) error {
	url := c.QueryParam("URL")
	if url == "" {
		return c.String(http.StatusBadRequest, "Invalid data")
	}

	jsUrl := "/js?url=" + url
	cssUrl := "/css?url=" + url

	data := IntroVideoData{
		Js:  jsUrl,
		Css: cssUrl,
	}

	return c.JSON(http.StatusOK, data)
}

func ServeJavaScript(c echo.Context) error {
	url := c.QueryParam("url")
	if url == "" {
		return c.String(http.StatusBadRequest, "Invalid data")
	}
	c.Set("url", url)

	scriptErr, js := Script(c)
	if scriptErr != nil {
		return scriptErr
	}

	c.Response().Header().Set(echo.HeaderContentType, "application/javascript")
	return c.String(http.StatusOK, js)
}

func ServeCSS(c echo.Context) error {
	url := c.QueryParam("url")
	if url == "" {
		return c.String(http.StatusBadRequest, "Invalid data")
	}
	c.Set("url", url)

	stylesheetErr, css := Stylesheet(c)
	if stylesheetErr != nil {
		return stylesheetErr
	}

	c.Response().Header().Set(echo.HeaderContentType, "text/css")
	return c.String(http.StatusOK, css)
}
