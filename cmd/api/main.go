package main

import (
	"os/exec"
	"strings"

	"github.com/crocoder-dev/intro-video/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getVideoResolution(videoURL string) (string, error) {
	// Prepare the ffmpeg command to get resolution
	cmd := exec.Command("ffprobe", "-v", "error", "-select_streams", "v:0", "-show_entries", "stream=width,height", "-of", "csv=p=0", videoURL)

	// Run the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	// The output should be in the format "widthxheight"
	resolution := strings.TrimSpace(string(output))
	return resolution, nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/wtf", func(c echo.Context) error {
		videoURL := "https://cdn.dribbble.com/userupload/92566/file/original-53ad0460a2ad35860f2859f174d7a6f4.mov" // Replace with your actual video URL
		resolution, err := getVideoResolution(videoURL)
		if err != nil {
			return nil
		}
		return c.String(200, resolution)
	})

	e.GET("/script.js", handler.Script)

	e.GET("/style.css", handler.Stylesheet)

	e.Logger.Fatal(e.Start(":8080"))
}
