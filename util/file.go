package util

import (
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"os"
	"strings"
)

func GetEmptyJsonNode() *jsonvalue.V {
	return jsonvalue.NewObject()
}

func GetEmptyJsonArray() *jsonvalue.V {
	return jsonvalue.NewArray()
}

func FileOrPathExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func FileNameAndExt(fileName string) (string, string) {
	splited := strings.Split(fileName, ".")
	if len(splited) < 2 {
		return fileName, ""
	}
	return strings.Join(splited[0:len(splited)-1], "."), splited[len(splited)-1]
}
