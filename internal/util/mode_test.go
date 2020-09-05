package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMode(t *testing.T) {
	list := map[string]Mode{
		"block": Mode("block"),
		"mute":  Mode("mute"),
		"test":  Mode("unknown"),
	}

	for k, v := range list {
		actual := NewMode(k)
		assert.Equal(t, v, actual)
	}
}

func TestMode_IsBlockMode(t *testing.T) {
	list := map[Mode]bool{
		NewMode("block"):   true,
		NewMode("mute"):    false,
		NewMode("unknown"): false,
	}

	for k, v := range list {
		actual := k.IsBlockMode()
		assert.Equal(t, v, actual)
	}
}

func TestMode_String(t *testing.T) {
	list := map[Mode]string{
		NewMode("block"):   "block",
		NewMode("mute"):    "mute",
		NewMode("unknown"): "unknown",
	}

	for k, v := range list {
		actual := k.String()
		assert.Equal(t, v, actual)
	}
}
