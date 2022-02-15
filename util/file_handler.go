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

var (
	JsonCommentReg, _  = regexp.Compile("//.*")
	JsonCommentReplace = []byte("")
)

var _fileHandler = make(map[string]FileHandlerFn)

func RegisterFileHandler(fileExt string, fn FileHandlerFn) {
	_fileHandler[fileExt] = fn
}

func GetFileHandler(fileExt string) FileHandlerFn {
	if _, exist := _fileHandler[fileExt]; !exist {
		return UnknownFileHandler
	}
	return _fileHandler[fileExt]
}

type FileHandlerFn func(bs []byte) (interface{}, error)

func (fn FileHandlerFn) Handle(bs []byte) (interface{}, error) {
	return fn(bs)
}

func JsonFileHandler(bs []byte) (interface{}, error) {
	rep := JsonCommentReg.ReplaceAll(bs, JsonCommentReplace)
	return sonic.Get(rep)
}

func UnknownFileHandler(bs []byte) (interface{}, error) {
	return nil, nil
}
