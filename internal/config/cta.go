package config

import "errors"

type CtaButtonType string

const (
	DefaultCta CtaButtonType = "default"
	CustomCta  CtaButtonType = "custom"
)

func NewCtaButtonType(ctaButtonType string) (CtaButtonType, error) {
	switch CtaButtonType(ctaButtonType) {
	case DefaultCta, CustomCta:
		return CtaButtonType(ctaButtonType), nil
	default:
		return "", errors.New("invalid CtaButtonType")
	}
}

type Cta struct {
	Enabled     bool
	TextContent string
	Type        CtaButtonType
}
