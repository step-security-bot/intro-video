package script

import (
	"testing"
)

func TestProcess(t *testing.T) {
	_, err := process(ScriptProps{
	})
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
