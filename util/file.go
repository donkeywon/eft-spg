package util

import (
	"github.com/bytedance/sonic/ast"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func ReadJsonDir(dir string) (ast.Node, error) {
	var m ast.Node

	fs, _ := ioutil.ReadDir(dir)
	for _, f := range fs {
		path := filepath.Join(dir, f.Name())
		fn, fe := FileNameAndExt(f.Name())
		if f.IsDir() {
			n, err := ReadJsonDir(path)
			if err != nil {
				return m, err
			}

			_, err = m.Set(fn, n)
			if err != nil {
				return m, err
			}

		} else {
			if fe != "json" {
				continue
			}

			fbs, err := ioutil.ReadFile(path)
			if err != nil {
				return m, err
			}

			n, err := GetFileHandler("json").Handle(fbs)
			if err != nil {
				return m, err
			}

			_, err = m.Set(fn, n.(ast.Node))
			if err != nil {
				return m, err
			}
		}
	}

	return m, nil
}

func FileNameAndExt(fileName string) (string, string) {
	splited := strings.Split(fileName, ".")
	if len(splited) < 2 {
		return fileName, ""
	}
	return strings.Join(splited[0:len(splited)-1], "."), splited[len(splited)-1]
}
