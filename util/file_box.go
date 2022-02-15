package util

import (
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/donkeywon/gtil/util"
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr/v2"
	"strings"
)

var (
	AssetsBox   = packr.New("Assets", "../assets")
	DatabaseBox = packr.New("Database", "../assets/database")
	ImageBox    = packr.New("Image", "../assets/images")
	ConfigBox   = packr.New("Config", "../cfg")

	PathSeparator = "/"
	EmptyJsonNode = []byte("{}")
)

func ReadJsonBox(box *packr.Box) (*ast.Node, error) {
	if box == nil {
		return &ast.Node{}, nil
	}

	n, _ := sonic.Get(EmptyJsonNode)

	err := box.Walk(func(filePath string, fileInfo packd.File) error {
		filePathSplit := strings.Split(filePath, PathSeparator)
		fn, fe := FileNameAndExt(filePathSplit[len(filePathSplit)-1])
		if fe != JsonFileExt {
			return nil
		}

		n1 := &n
		for i := 0; i < len(filePathSplit)-1; i++ {
			if !n1.Get(filePathSplit[i]).Exists() {
				newN, _ := sonic.Get(EmptyJsonNode)
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
		_, err = n1.SetAny(fn, n2.(ast.Node))
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

func ReadDatabaseBox() (*ast.Node, error) {
	return ReadJsonBox(DatabaseBox)
}
