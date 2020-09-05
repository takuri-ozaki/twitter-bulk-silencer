package user

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
	"twitter-bulk-silencer/internal/util"
)

func TestHandler_SaveUserList(t *testing.T) {
	list := map[string]string{
		"block":    "1\n2\n3\n4\n5\n",
		"mute":     "11\n12\n13\n14\n15\n",
		"follower": "21\n22\n23\n24\n25\n",
		"followee": "31\n32\n33\n34\n35\n",
	}

	for k, v := range list {
		handler := NewHandler(util.NewDummyAPI(), util.NewTarget(k), util.NewDummyFileSystem(map[string][]byte{k + "list.txt": {}}))
		err := handler.SaveUserList()
		assert.Nil(t, err)
		buffer, _ := handler.FileSystem.OpenFile(k+"list.txt", false)
		byt, _ := ioutil.ReadAll(buffer)
		assert.Equal(t, v, string(byt))
	}
}

func TestHandler_WriteAppendUserList(t *testing.T) {
	handler := NewHandler(util.NewDummyAPI(), util.NewTarget("unittest"), util.NewDummyFileSystem(map[string][]byte{"unknownlist.txt": []byte("41\n42\n43\n44\n45\n")}))
	err := handler.WriteAppendUserList([]int64{1, 2, 3, 4, 5})
	assert.Nil(t, err)
	buffer, _ := handler.FileSystem.OpenFile("unknownlist.txt", false)
	byt, _ := ioutil.ReadAll(buffer)
	assert.Equal(t, "41\n42\n43\n44\n45\n1\n2\n3\n4\n5\n", string(byt))
}

func TestHandler_ReadUserList(t *testing.T) {
	handler := NewHandler(util.NewDummyAPI(), util.NewTarget("unittest"), util.NewDummyFileSystem(map[string][]byte{"unknownlist.txt": []byte("41\n42\n43\n44\n45\n")}))
	list, err := handler.ReadUserList()
	assert.Nil(t, err)
	assert.Equal(t, []int64{41, 42, 43, 44, 45}, list)
}
