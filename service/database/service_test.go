package database

import (
	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSvc_Open(t *testing.T) {
	path := "/Users/donkeywon/code/go/eft-spg/assets/database"
	bs := make([]byte, 500000000, 500000000)
	n, err := readDirToJson(bs, path)
	assert.NoError(t, err)

	_, err = sonic.Get(bs[:n])
	assert.NoError(t, err)
	bs = nil

	time.Sleep(time.Second * 100)
}

//func TestSvc_Open(t *testing.T) {
//	n := util.GetEmptyJsonNode()
//
//	start := time.Now().UnixNano()
//	filepath.Walk("assets/database", func(path string, info fs.FileInfo, err error) error {
//		if info.IsDir() {
//			return nil
//		}
//
//		filePathSplit := strings.Split(path, string(os.PathSeparator))
//		fn, fe := util.FileNameAndExt(filePathSplit[len(filePathSplit)-1])
//		if fe != util.JsonFileExt {
//			return nil
//		}
//		filePathSplit[len(filePathSplit)-1] = fn
//
//		bs, err := ioutil.ReadFile(path)
//		if err != nil {
//			return errors.Wrapf(err, util.ErrReadFileBox, path)
//		}
//
//		bs, err = util.GetFileHandler(fe).Handle(bs)
//		if err != nil {
//			return errors.Wrapf(err, util.ErrReadFileBox, path)
//		}
//
//		i := skipBlank(bs)
//		var n1 interface{}
//		if bs[i] == '[' {
//			n1 = util.GetEmptyJsonArray()
//		} else {
//			n1 = util.GetEmptyJsonNode()
//		}
//		jsoniter.Unmarshal(bs, &n1)
//
//		n.Set(n1, filePathSplit[2:]...)
//
//		return nil
//	})
//	end := time.Now().UnixNano()
//	fmt.Println("read: ", (end-start)/1e6)
//
//	start = time.Now().UnixNano()
//
//	time.Sleep(time.Second * 100)
//
//	//node, _, _ := n.Get("locations", "rezervbase", "loot")
//	//jsonparser.ObjectEach(node, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
//	//	fmt.Println(string(key))
//	//	return nil
//	//})
//
//	//for i := 0; i < 10; i++ {
//	//	n.Get("locations", "rezervbase", "loot", "dynamic")
//	//	//jsonparser.GetInt(bs, "locations", "woods", "base", "sav_summon_seconds")
//	//	//jsoniter.Get(bs, "locations", "woods", "base", "sav_summon_seconds").ToInt()
//	//}
//	//end = time.Now().UnixNano()
//	//fmt.Println("find: ", (end-start)/1e6)
//}
//
//func skipBlank(bs []byte) int {
//	for i := 0; i < len(bs); i++ {
//		if bs[i] != '\n' && bs[i] != ' ' && bs[i] != '\t' {
//			return i
//		}
//	}
//
//	return -1
//}

//import (
//	"eft-spg/util"
//	"fmt"
//	"github.com/buger/jsonparser"
//	util2 "github.com/donkeywon/gtil/util"
//	"github.com/gobuffalo/packd"
//	"github.com/pkg/errors"
//	"github.com/stretchr/testify/assert"
//	"os"
//	"strings"
//)
//
//type JsonNode map[string]interface{}
//
//func NewJsonNode() JsonNode {
//	return make(JsonNode)
//}
//
//func (jn JsonNode) Set(bs []byte, paths ...string) {
//	if bs == nil {
//		return
//	}
//
//	if len(paths) == 0 {
//		return
//	}
//
//	i := 0
//	jn1 := jn
//	for {
//		if i == len(paths)-1 {
//			jn1[paths[i]] = bs
//			break
//		}
//
//		if _, exist := jn1[paths[i]]; !exist {
//			jn1[paths[i]] = NewJsonNode()
//		}
//		jn1 = jn1[paths[i]].(JsonNode)
//
//		i++
//	}
//}
//
//func (jn JsonNode) GetNodeByPath(path string, paths ...string) (JsonNode, error) {
//	jn1 := jn[path]
//	var exist bool
//	for _, p := range paths {
//		if _, ok := jn1.(JsonNode); !ok {
//			return nil, errors.New("Node not exist")
//		}
//
//		jn1, exist = jn1.(JsonNode)[p]
//		if !exist {
//			return nil, errors.New("Node not exist")
//		}
//	}
//
//	return jn1.(JsonNode), nil
//}
//
//func (jn JsonNode) Get(path string, paths ...string) ([]byte, jsonparser.ValueType, error) {
//	jn1 := jn[path]
//	for i, p := range paths {
//		if _, ok := jn1.(JsonNode); ok {
//			jn1 = jn1.(JsonNode)[p]
//			continue
//		}
//
//		if bs, ok := jn1.([]byte); ok {
//			v, vt, _, err := jsonparser.Get(bs, paths[i:]...)
//			return v, vt, err
//		}
//	}
//
//	if _, ok := jn1.([]byte); !ok {
//		return nil, jsonparser.NotExist, errors.New("Node not exist")
//	}
//
//	return jn1.([]byte), jsonparser.Unknown, nil
//}
//
//func TestService_Open(t *testing.T) {
//	box := util.DatabaseBox
//	n := NewJsonNode()
//
//	err := box.Walk(func(filePath string, fileInfo packd.File) error {
//		filePathSplit := strings.Split(filePath, string(os.PathSeparator))
//		fn, fe := util.FileNameAndExt(filePathSplit[len(filePathSplit)-1])
//		if fe != util.JsonFileExt {
//			return nil
//		}
//		filePathSplit[len(filePathSplit)-1] = fn
//
//		bs, err := util.GetFileHandler(fe).Handle(util2.String2Bytes(fileInfo.String()))
//		if err != nil {
//			return errors.Wrapf(err, util.ErrReadFileBox, filePath)
//		}
//
//		//v, err := jsonvalue.Unmarshal(bs)
//		//if err != nil {
//		//	return errors.Wrapf(err, util.ErrReadFileBox, filePath)
//		//}
//
//		n.Set(bs, filePathSplit...)
//		//
//		//if len(filePathSplit) == 1 {
//		//	n.Set(bs, filePathSplit[0])
//		//} else {
//		//	n.Set()
//		//	var at []interface{}
//		//	for _, p := range filePathSplit[1:] {
//		//		at = append(at, p)
//		//	}
//		//
//		//	_, err = n.Set(v).At(filePathSplit[0], at...)
//		//}
//		//if err != nil {
//		//	return errors.Wrapf(err, util.ErrReadFileBox, filePath)
//		//}
//
//		return nil
//	})
//
//	assert.NoError(t, err, "read box fail")
//	fmt.Println(n.Get("bots", "types", "bossgluhar", "appearance", "body"))
//	fmt.Println(n.Get("bots"))
//	_, err = n.GetNodeByPath("bots")
//	assert.NoError(t, err)
//	_, err = n.GetNodeByPath("bots", "types")
//	assert.NoError(t, err)
//	_, err = n.GetNodeByPath("nots", "types", "bossgluhar")
//	assert.Error(t, err)
//}
