package util

import (
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

type FileHandlerFn func(bs []byte) ([]byte, error)

func (fn FileHandlerFn) Handle(bs []byte) ([]byte, error) {
	return fn(bs)
}

func JsonFileHandler(bs []byte) ([]byte, error) {
	rep := JsonCommentReg.ReplaceAll(bs, JsonCommentReplace)
	return rep, nil
}

func UnknownFileHandler(bs []byte) ([]byte, error) {
	return nil, nil
}
