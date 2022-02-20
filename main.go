package main

import (
	"bytes"
	"eft-spg/util"
	"fmt"
	"github.com/pkg/errors"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {

	bs := &bytes.Buffer{}
	filepath.Walk("assets/database", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		filePathSplit := strings.Split(path, string(os.PathSeparator))
		fn, fe := util.FileNameAndExt(filePathSplit[len(filePathSplit)-1])
		if fe != util.JsonFileExt {
			return nil
		}
		filePathSplit[len(filePathSplit)-1] = fn

		bs1, err := ioutil.ReadFile(path)
		if err != nil {
			return errors.Wrapf(err, util.ErrReadFileBox, path)
		}

		bs.Write(bs1)

		return nil
	})

	fmt.Println("read done: len: ", bs.Len())
	time.Sleep(time.Second * 100)

	//bs, err := readDirToJson("assets/database")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println("read done, len: ", len(bs))
	//
	////pj, err := simdjson.Parse(bs, nil)
	////if err != nil {
	////   fmt.Println(err)
	////}
	////_ = pj
	////fmt.Println("parse done")
	//
	////var m map[string]interface{}
	////err = jsoniter.Unmarshal(bs, &m)
	////if err != nil {
	////	fmt.Println(err)
	////}
	////fmt.Println("parse done")
	//bs = nil
	//
	//time.Sleep(time.Second * 100)

	//l, _ := logger.FromConfig(logger.DefaultConsoleConfig(), zap.WrapCore(core.NewStackExtractCore))
	//zap.ReplaceGlobals(l)
	//
	//config := cmd.NewConfig()
	//f, err := ioutil.ReadFile("./config.yaml")
	//if err != nil {
	//	l.Error("Start fail", zap.Error(err))
	//	return
	//}
	//
	//err = yaml.Unmarshal(f, config)
	//if err != nil {
	//	l.Error("Start fail", zap.Error(err))
	//	return
	//}
	//
	//c := cmd.New(config)
	//err = service.DoOpen(c, context.Background(), l)
	//if err != nil {
	//	l.Error("Start fail", zap.Error(err))
	//	return
	//}
	//
	//l.Info("Start success")
	//
	//signalCh := make(chan os.Signal)
	//signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	//
	//<-signalCh
	//l.Info("Receive signal, exit")
	//
	//go func() {
	//	err := service.DoClose(c)
	//	if err != nil {
	//		l.Error("Close fail", zap.Error(err))
	//	}
	//}()
	//
	//select {
	//case <-c.Closed():
	//	l.Info("Closed")
	//case <-signalCh:
	//	l.Info("Receive signal twice, exit")
	//case <-time.After(time.Second * 10):
	//	l.Info("Close timeout, exit")
	//}
}

func readDirToJson(path string) ([]byte, error) {
	if util.FileExist(path) {
		return ioutil.ReadFile(path)
	}

	if !util.DirExist(path) {
		return nil, errors.New("Path not exist")
	}

	infos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	bs := []byte{'{'}
	//bs := &bytes.Buffer{}
	//bs.WriteByte('{')
	for _, info := range infos {
		fn, fe := util.FileNameAndExt(info.Name())
		if fe != util.JsonFileExt && util.FileExist(filepath.Join(path, info.Name())) {
			continue
		}

		bs = append(bs, []byte(`"`+fn+`":`)...)
		//bs.WriteString(`"` + fn + `":`)
		bs1, err := readDirToJson(filepath.Join(path, info.Name()))
		if err != nil {
			return nil, err
		}

		bs = append(bs, bs1...)
		bs = append(bs, ',')
		//bs.Write(bs1)
		//bs.WriteByte(',')
	}
	if bs[len(bs)-1] == ',' {
		bs = bs[:len(bs)-1]
	}
	bs = append(bs, '}')
	//if bs.Bytes()[bs.Len()-1] == ',' {
	//    bs.Truncate(bs.Len() - 1)
	//}
	//bs.WriteByte('}')

	return bs, nil
}
