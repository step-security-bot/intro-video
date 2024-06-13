package config

import "errors"

type BubbleType string

const (
	DefaultBubble BubbleType = "default"
	CustomBubble  BubbleType = "custom"
)

func NewBubbleType(bubbleType string) (BubbleType, error) {
	switch BubbleType(bubbleType) {
	case DefaultBubble, CustomBubble:
		return BubbleType(bubbleType), nil
	default:
		return "", errors.New("invalid BubbleType")
	}
}

type Bubble struct {
	Enabled     bool
	TextContent string
	Type        BubbleType
}
