package config

type CtaButtonType int

const (
    DefaultCta CtaButtonType = iota
	CustomCta
)


type Cta struct {
	Enabled bool
	TextContent string
	Type CtaButtonType
}
