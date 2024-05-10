package script

import (
	"bytes"
	"html/template"
	"io"
	"os"

	"github.com/tdewolff/minify/v2/minify"
)

type ScriptProps struct {
	Bubble
	Cta
}

type Bubble struct {
	Enabled bool
	TextContent string
}

type Cta struct {
	Enabled bool
	TextContent string
}

func process(props ScriptProps) (interface{}, error) {

	t, err := template.ParseFiles(
		"../template/script/bubble.js.tmpl",
		"../template/script/cta.js.tmpl",
	)

	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer

	err = t.ExecuteTemplate(&buf, "bubble", props.Bubble)
	if err != nil {
		return nil, err
	}

	err = t.ExecuteTemplate(&buf, "cta", props.Cta)
	if err != nil {
		return nil, err
	}

	file, err := os.Open("../template/script/base.js")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	base, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var result bytes.Buffer
	result.Write(buf.Bytes())
	result.Write(base)

	m := minify.Default
	out, err := m.String("text/javascript", result.String())
	if err != nil {
		return nil, err
	}

	return out, nil
}
