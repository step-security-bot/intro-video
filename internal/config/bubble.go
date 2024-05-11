package config

type BubbleType int

const (
	DefaultBubble BubbleType = iota
)


type Bubble struct {
	Enabled bool
	TextContent string
	Type BubbleType
}

