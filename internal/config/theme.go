package config

import "errors"

type Theme string

const (
	DefaultTheme Theme = "default"
)

func NewTheme(theme string) (Theme, error) {
	switch Theme(theme) {
	case DefaultTheme:
		return Theme(theme), nil
	default:
		return "", errors.New("invalid Theme")
	}
}
