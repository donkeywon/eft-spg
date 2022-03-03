package util

import (
	"github.com/bytedance/sonic/ast"
	"github.com/donkeywon/gtil/util"
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr/v2"
	"strings"
)

var (
	ConfigBox = packr.New("Config", "../cfg")
)

var (
	PathSeparator = "/"
)

func ReadJsonBox(box *packr.Box) (*ast.Node, error) {
	n := GetEmptyJsonNode()
	if box == nil {
		return &n, nil
	}

	err := box.Walk(func(filePath string, fileInfo packd.File) error {
		filePathSplit := strings.Split(filePath, PathSeparator)
		fn, fe := FileNameAndExt(filePathSplit[len(filePathSplit)-1])
		if fe != JsonFileExt {
			return nil
		}

		n1 := &n
		for i := 0; i < len(filePathSplit)-1; i++ {
			if !n1.Get(filePathSplit[i]).Exists() {
				newN := GetEmptyJsonNode()
				_, err := n1.Set(filePathSplit[i], newN)
				if err != nil {
					return err
				}
			}

			n1 = n1.Get(filePathSplit[i])
		}

		n2, err := GetFileHandler(fe).Handle(util.String2Bytes(fileInfo.String()))
		if err != nil {
			return err
		}
		_, err = n1.Set(fn, n2.(ast.Node))
		if err != nil {
			return err
		}

		return nil
	})

	return &n, err
}

func ReadConfigBox() (*ast.Node, error) {
	return ReadJsonBox(ConfigBox)
}
