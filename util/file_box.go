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
	AssetsFileBox   = packr.New("Assets", "../assets")
	DatabaseFileBox = packr.New("Database", "../assets/database")
	ImageFileBox    = packr.New("Image", "../assets/images")
	ConfigFileBox   = packr.New("Config", "../config")

	PathSeparator = "/"
)

func ReadJsonBox(box *packr.Box) (*ast.Node, error) {
	if box == nil {
		return &ast.Node{}, nil
	}

	n, err := sonic.Get([]byte("{}"))
	if err != nil {
		return nil, err
	}

	box.Walk(func(filePath string, fileInfo packd.File) error {
		filePathSplit := strings.Split(filePath, PathSeparator)
		n1 := &n
		for i := 0; i < len(filePathSplit)-1; i++ {
			if !n1.Get(filePathSplit[i]).Exists() {
				_, err := n1.Set(filePathSplit[i], ast.NewNull())
				if err != nil {
					return err
				}

				n1 = n1.Get(filePathSplit[i])
			}
		}

		fn, fe := FileNameAndExt(filePathSplit[len(filePathSplit)-1])
		_, err := GetFileHandler(fe).Handle(util.String2Bytes(fileInfo.String()))
		if err != nil {
			return err
		}
		_, err = n1.Set(fn, ast.NewNull())
		if err != nil {
			return err
		}

		return nil
	})

	return &n, nil
}
