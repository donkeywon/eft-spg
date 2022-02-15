package util

import (
	"github.com/buger/jsonparser"
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

	EmptyJsonNode = []byte("{}")
	PathSeparator = "/"
)

func ReadJsonBox(box *packr.Box) ([]byte, error) {
	n := EmptyJsonNode
	if box == nil {
		return n, nil
	}

	err := box.Walk(func(filePath string, fileInfo packd.File) error {
		filePathSplit := strings.Split(filePath, PathSeparator)
		fn, fe := FileNameAndExt(filePathSplit[len(filePathSplit)-1])
		if fe != JsonFileExt {
			return nil
		}
		filePathSplit[len(filePathSplit)-1] = fn

		n2, err := GetFileHandler(fe).Handle(util.String2Bytes(fileInfo.String()))
		if err != nil {
			return err
		}
		n, err = jsonparser.Set(n, n2, filePathSplit...)
		if err != nil {
			return err
		}

		return nil
	})

	return n, err
}

func ReadConfigBox() ([]byte, error) {
	return ReadJsonBox(ConfigBox)
}

func ReadDatabaseBox() ([]byte, error) {
	return ReadJsonBox(DatabaseBox)
}

func FileNameAndExt(fileName string) (string, string) {
	splited := strings.Split(fileName, ".")
	if len(splited) < 2 {
		return fileName, ""
	}
	return strings.Join(splited[0:len(splited)-1], "."), splited[len(splited)-1]
}
