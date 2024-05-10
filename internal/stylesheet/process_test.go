package stylesheet

import (
	"testing"
)

func TestProcess(t *testing.T) {
	_, err := process(StylesheetProps{
		Dimensions{100, 200, 300, 400},
		Bubble{},
		Cta{},
	})
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
