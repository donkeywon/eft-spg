package util

import (
	"github.com/buger/jsonparser"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

type JsonNode map[string]interface{}

type JsonArray []interface{}

func NewJsonNode() JsonNode {
	return make(JsonNode)
}

func NewJsonArray() JsonArray {
	return make(JsonArray, 0)
}

func (jn JsonNode) Set(n interface{}, paths ...string) {
	if n == nil {
		return
	}

	if len(paths) == 0 {
		return
	}

	i := 0
	jn1 := jn
	for {
		if i == len(paths)-1 {
			jn1[paths[i]] = n
			break
		}

		if _, exist := jn1[paths[i]]; !exist {
			jn1[paths[i]] = NewJsonNode()
		}
		jn1 = jn1[paths[i]].(JsonNode)

		i++
	}
}

func (jn JsonNode) Marshal() ([]byte, error) {
	// jsoniter 4340
	// go-json   7000
	return jsoniter.Marshal(jn)
}

func (jn JsonNode) GetNodeByPath(path string, paths ...string) (JsonNode, error) {
	jn1 := jn[path]
	var exist bool
	for _, p := range paths {
		if _, ok := jn1.(JsonNode); !ok {
			return nil, errors.New("Node not exist")
		}

		jn1, exist = jn1.(JsonNode)[p]
		if !exist {
			return nil, errors.New("Node not exist")
		}
	}

	return jn1.(JsonNode), nil
}

func (jn JsonNode) Get(path string, paths ...string) ([]byte, jsonparser.ValueType, error) {
	jn1 := jn[path]
	for i, p := range paths {
		if _, ok := jn1.(JsonNode); ok {
			jn1 = jn1.(JsonNode)[p]
			continue
		}

		if bs, ok := jn1.([]byte); ok {
			v, vt, _, err := jsonparser.Get(bs, paths[i:]...)
			return v, vt, err
		}
	}

	if _, ok := jn1.([]byte); !ok {
		return nil, jsonparser.NotExist, errors.New("Node not exist")
	}

	return jn1.([]byte), jsonparser.Unknown, nil
}
