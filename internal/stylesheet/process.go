package stylesheet

import (
	"bytes"
	"io"
	"os"
	"text/template"

	"github.com/tdewolff/minify/v2/minify"
)

type StylesheetProps struct {
	Dimensions
	Bubble
	Cta
}

type Dimensions struct {
	Swidth  int
	Sheight int
	Lwidth  int
	Lheight int
}

type Bubble struct {
}

type Cta struct {
}



func process(props StylesheetProps) (interface{}, error) {

	t, err := template.ParseFiles(
		"../template/stylesheet/bubble.css.tmpl",
		"../template/stylesheet/cta.css.tmpl",
		"../template/stylesheet/dimension.css.tmpl",
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

	err = t.ExecuteTemplate(&buf, "dimension", props.Dimensions)
	if err != nil {
		return nil, err
	}

	file, err := os.Open("../template/stylesheet/base.css")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	base, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var result bytes.Buffer
	result.Write(base)
	result.Write(buf.Bytes())

	m := minify.Default
	out, err := m.String("text/css", result.String())
	if err != nil {
		return nil, err
	}

	return out, nil
}
