package util

import (
	"github.com/bytedance/sonic"
	"regexp"
)

func init() {
	RegisterFileHandler(JsonFileExt, JsonFileHandler)
}

const (
	JsonFileExt = "json"
)

var _fileHandler = make(map[string]FileHandlerFn)

func RegisterFileHandler(fileExt string, fn FileHandlerFn) {
	_fileHandler[fileExt] = fn
}

func GetFileHandler(fileExt string) FileHandlerFn {
	return _fileHandler[fileExt]
}

type FileHandlerFn func(bs []byte) (interface{}, error)

func (fn FileHandlerFn) Handle(bs []byte) (interface{}, error) {
	return fn(bs)
}

func JsonFileHandler(bs []byte) (interface{}, error) {
	re, _ := regexp.Compile("//.*")
	rep := re.ReplaceAll(bs, []byte{})
	return sonic.Get(rep)
}
