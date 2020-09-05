package util

import (
	"bytes"
	"io"
	"os"
)

type FileSystem interface {
	OpenFile(name string, append bool) (io.ReadWriteCloser, error)
}

type RealFileSystem struct {
	baseDir string
}

func NewRealFileSystem(baseDir string) FileSystem {
	return RealFileSystem{baseDir: baseDir}
}

func (r RealFileSystem) OpenFile(name string, append bool) (io.ReadWriteCloser, error) {
	if append {
		return os.OpenFile(r.baseDir+name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	}
	return os.OpenFile(r.baseDir+name, os.O_RDWR|os.O_CREATE, 0666)
}

type DummyFileSystem struct {
	DummyFileMap *map[string]DummyFile
}

func NewDummyFileSystem(buffers map[string][]byte) FileSystem {
	dummyFileMap := map[string]DummyFile{}
	for k, v := range buffers {
		dummyFileMap[k] = newDummyFile(v)
	}

	return DummyFileSystem{DummyFileMap: &dummyFileMap}
}

func (d DummyFileSystem) OpenFile(name string, _ bool) (io.ReadWriteCloser, error) {
	mp := *d.DummyFileMap
	return mp[name], nil
}

type DummyFile struct {
	*bytes.Buffer
}

func newDummyFile(buffer []byte) DummyFile {
	return DummyFile{bytes.NewBuffer(buffer)}
}

func (d DummyFile) Close() error {
	return nil
}
