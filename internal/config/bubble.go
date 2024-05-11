package config

type BubbleType string

const (
	DefaultBubble BubbleType = "default"
	CustomBubble BubbleType = "custom"
)

type Bubble struct {
	Enabled bool
	TextContent string
	Type BubbleType
}

