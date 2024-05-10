package script

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/tdewolff/minify/v2/minify"
)

type ScriptProps struct {
}


func process(props ScriptProps) (interface{}, error) {

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
	result.Write(base)

	m := minify.Default
	out, err := m.String("text/javascript", result.String())
	if err != nil {
		return nil, err
	}

	fmt.Println(out)

	return out, nil
}
