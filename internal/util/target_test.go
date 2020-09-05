package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTarget(t *testing.T) {
	list := map[string]RealTarget{
		"block":    RealTarget("block"),
		"mute":     RealTarget("mute"),
		"follower": RealTarget("follower"),
		"followee": RealTarget("followee"),
		"test":     RealTarget("unknown"),
	}

	for k, v := range list {
		actual := NewTarget(k)
		assert.Equal(t, v, actual)
	}
}

func TestTarget_GetFilePath(t *testing.T) {
	list := map[RealTarget]string{
		NewTarget("block"):    "blocklist.txt",
		NewTarget("mute"):     "mutelist.txt",
		NewTarget("follower"): "followerlist.txt",
		NewTarget("followee"): "followeelist.txt",
	}

	for k, v := range list {
		actual := k.GetFileName()
		assert.Equal(t, v, actual)
	}
}

func TestTarget_String(t *testing.T) {
	list := map[RealTarget]string{
		NewTarget("block"):    "block",
		NewTarget("mute"):     "mute",
		NewTarget("follower"): "follower",
		NewTarget("followee"): "followee",
		NewTarget("unknown"):  "unknown",
	}

	for k, v := range list {
		actual := k.String()
		assert.Equal(t, v, actual)
	}
}
