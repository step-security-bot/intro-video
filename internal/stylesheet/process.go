package stylesheet

import (
	"bytes"
	"io"
	"os"
	"text/template"

	"github.com/tdewolff/minify/v2/minify"
)

type StylesheetProps struct {
	Bubble
	Cta
}

type Bubble struct {
}

type Cta struct {
}



func Process(props StylesheetProps) (string, error) {

	t, err := template.ParseFiles(
		"internal/template/stylesheet/bubble.css.tmpl",
		"internal/template/stylesheet/cta.css.tmpl",
	)

	if err != nil {
		return "", err
	}
	var buf bytes.Buffer

	err = t.ExecuteTemplate(&buf, "bubble", props.Bubble)
	if err != nil {
		return "", err
	}

	err = t.ExecuteTemplate(&buf, "cta", props.Cta)
	if err != nil {
		return "", err
	}

	file, err := os.Open("internal/template/stylesheet/base.css")
	if err != nil {
		return "", err
	}
	defer file.Close()

	base, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	var result bytes.Buffer
	result.Write(base)
	result.Write(buf.Bytes())

	m := minify.Default
	out, err := m.String("text/css", result.String())
	if err != nil {
		return "", err
	}

	return out, nil
}
