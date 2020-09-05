package silence

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
	"twitter-bulk-silencer/internal/util"
)

var buffers = map[string][]byte{
	"blocklist.txt":    []byte(""),
	"mutelist.txt":     []byte(""),
	"followerlist.txt": []byte("21\n22\n23\n24\n25\n"),
	"followeelist.txt": []byte("31\n32\n33\n34\n35\n"),
}

var idList = []int64{1001, 1002, 1003, 1004, 1005, 1006, 3, 13, 23, 33}

type input struct {
	protectFollower bool
	protectFollowee bool
	execute         bool
	expected        string
}

func TestHandler_Silence(t *testing.T) {
	list := []input{
		{true, true, true, "1001\n1002\n1003\n1004\n1005\n1006\n3\n13\n"},
		{false, true, true, "1001\n1002\n1003\n1004\n1005\n1006\n3\n13\n23\n"},
		{true, false, true, "1001\n1002\n1003\n1004\n1005\n1006\n3\n13\n33\n"},
		{false, false, true, "1001\n1002\n1003\n1004\n1005\n1006\n3\n13\n23\n33\n"},
		{false, false, false, ""},
	}

	for _, v := range list {
		handler := NewHandler(
			util.NewDummyAPI(),
			true,
			5,
			v.protectFollower,
			v.protectFollowee,
			v.execute,
			util.NewDummyFileSystem(buffers),
		)

		err := handler.Silence(idList)
		assert.Nil(t, err)

		buffer, _ := handler.silenceeHandler.FileSystem.OpenFile("blocklist.txt", false)
		byt, _ := ioutil.ReadAll(buffer)
		assert.Equal(t, v.expected, string(byt))
	}
}
