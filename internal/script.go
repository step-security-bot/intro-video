package internal

import (
	"bytes"
	"errors"
	"html/template"
	"io"
	"os"

	"github.com/tdewolff/minify/v2/minify"
)

type Script struct {}

func (s Script) Process(props ProcessableFileProps) (string, error) {
	if props.Video.URL == "" {
		return "", errors.New("video URL is required")
	}

	t, err := template.ParseFiles(
		"internal/template/script/start.js.tmpl",
		"internal/template/script/end.js.tmpl",
		"internal/template/script/video.js.tmpl",
		"internal/template/script/bubble.js.tmpl",
		"internal/template/script/cta.js.tmpl",
	)

	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	var end bytes.Buffer

	err = t.ExecuteTemplate(&buf, "start", props.Bubble)
	if err != nil {
		return "", err
	}

	err = t.ExecuteTemplate(&buf, "video", props.Video)
	if err != nil {
		return "", err
	}

	if props.Bubble.Enabled {
		err = t.ExecuteTemplate(&buf, "bubble", props.Bubble)
		if err != nil {
			return "", err
		}
	}

	if props.Cta.Enabled {
		err = t.ExecuteTemplate(&buf, "cta", props.Cta)
		if err != nil {
			return "", err
		}
	}

	err = t.ExecuteTemplate(&end, "end", nil)
	if err != nil {
		return "", err
	}

	file, err := os.Open("internal/template/script/base.js")
	if err != nil {
		return "", err
	}
	defer file.Close()

	base, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	var result bytes.Buffer
	result.Write(buf.Bytes())
	result.Write(base)
	result.Write(end.Bytes())

	m := minify.Default
	out, err := m.String("text/javascript", result.String())
	if err != nil {
		return "", err
	}

	return out, nil
}