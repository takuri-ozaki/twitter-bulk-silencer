package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"twitter-bulk-silencer/internal/util"
)

func TestHandler_GetTargetUsers(t *testing.T) {
	handler := NewHandler(util.NewDummyAPI(), "test")
	list, err := handler.GetTargetUsers()
	assert.Nil(t, err)
	assert.Equal(t, []int64{1001, 1002, 1003, 1004, 1005, 1006, 3, 13, 23, 33}, list)
}
