package config

type CtaButtonType int

const (
    DefaultCta CtaButtonType = iota
)


type Cta struct {
	Enabled bool
	TextContent string
	Type CtaButtonType
}
