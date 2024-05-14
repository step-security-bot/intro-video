package config

type CtaButtonType string

const (
	DefaultCta CtaButtonType = "default"
	CustomCta  CtaButtonType = "custom"
)

type Cta struct {
	Enabled     bool
	TextContent string
	Type        CtaButtonType
}
