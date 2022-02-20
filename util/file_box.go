package util

import (
	"github.com/gobuffalo/packr/v2"
)

var (
	AssetsBox   = packr.New("Assets", "../assets")
	DatabaseBox = packr.New("Database", "../assets/database")
	ImageBox    = packr.New("Image", "../assets/images")
	ConfigBox   = packr.New("Config", "../cfg")
)

var (
	PathSeparator = "/"
)

func ReadJsonBox(box *packr.Box) (JsonNode, error) {
	return nil, nil
	//n := GetEmptyJsonNode()
	//if box == nil {
	//	return n, nil
	//}
	//
	//err := box.Walk(func(filePath string, fileInfo packd.File) error {
	//	filePathSplit := strings.Split(filePath, PathSeparator)
	//	fn, fe := FileNameAndExt(filePathSplit[len(filePathSplit)-1])
	//	if fe != JsonFileExt {
	//		return nil
	//	}
	//	filePathSplit[len(filePathSplit)-1] = fn
	//
	//	bs, err := GetFileHandler(fe).Handle(util.String2Bytes(fileInfo.String()))
	//	if err != nil {
	//		return errors.Wrapf(err, ErrReadFileBox, filePath)
	//	}
	//
	//	v, err := jsonvalue.Unmarshal(bs)
	//	if err != nil {
	//		return errors.Wrapf(err, ErrReadFileBox, filePath)
	//	}
	//
	//	if len(filePathSplit) == 1 {
	//		_, err = n.Set(v).At(filePathSplit[0])
	//	} else {
	//		var at []interface{}
	//		for _, p := range filePathSplit[1:] {
	//			at = append(at, p)
	//		}
	//
	//		_, err = n.Set(v).At(filePathSplit[0], at...)
	//	}
	//	if err != nil {
	//		return errors.Wrapf(err, ErrReadFileBox, filePath)
	//	}
	//
	//	return nil
	//})
	//
	//return n, err
}

func ReadConfigBox() (JsonNode, error) {
	return ReadJsonBox(ConfigBox)
}

func ReadDatabaseBox() (JsonNode, error) {
	return ReadJsonBox(DatabaseBox)
}
