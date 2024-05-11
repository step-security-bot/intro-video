package config

type BubbleType int

const (
	DefaultBubble BubbleType = iota
	CustomBubble
)

type Bubble struct {
	Enabled bool
	TextContent string
	Type BubbleType
}

