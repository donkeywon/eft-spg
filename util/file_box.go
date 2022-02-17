package util

import (
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"github.com/donkeywon/gtil/util"
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr/v2"
	"github.com/pkg/errors"
	"strings"
)

var (
	AssetsBox   = packr.New("Assets", "../assets")
	DatabaseBox = packr.New("Database", "../assets/database")
	ImageBox    = packr.New("Image", "../assets/images")
	ConfigBox   = packr.New("Config", "../cfg")

	EmptyJsonNode = jsonvalue.MustUnmarshalString("{}")
	PathSeparator = "/"
)

func ReadJsonBox(box *packr.Box) (*jsonvalue.V, error) {
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

		bs, err := GetFileHandler(fe).Handle(util.String2Bytes(fileInfo.String()))
		if err != nil {
			return errors.WithMessagef(err, ErrReadFileBox, filePath)
		}

		v, err := jsonvalue.Unmarshal(bs)
		if err != nil {
			return errors.Wrapf(err, ErrReadFileBox, filePath)
		}

		if len(filePathSplit) == 1 {
			_, err = n.Set(v).At(filePathSplit[0])
		} else {
			var at []interface{}
			for _, p := range filePathSplit[1:] {
				at = append(at, p)
			}

			_, err = n.Set(v).At(filePathSplit[0], at...)
		}
		if err != nil {
			return errors.WithMessagef(err, ErrReadFileBox, filePath)
		}

		return nil
	})

	return n, err
}

func ReadConfigBox() (*jsonvalue.V, error) {
	return ReadJsonBox(ConfigBox)
}

func ReadDatabaseBox() (*jsonvalue.V, error) {
	return ReadJsonBox(DatabaseBox)
}

func FileNameAndExt(fileName string) (string, string) {
	splited := strings.Split(fileName, ".")
	if len(splited) < 2 {
		return fileName, ""
	}
	return strings.Join(splited[0:len(splited)-1], "."), splited[len(splited)-1]
}
