package util

import (
	"github.com/bytedance/sonic/ast"
	"os"
	"strings"
)

func GetEmptyJsonNode() ast.Node {
	return ast.NewObject([]ast.Pair{})
}

func GetEmptyJsonArray() ast.Node {
	return ast.NewArray([]ast.Node{})
}

func FileExist(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	if info.IsDir() {
		return false
	}

	return true
}

func DirExist(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	if !info.IsDir() {
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
