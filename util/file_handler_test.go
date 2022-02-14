package util

import (
	"fmt"
	"github.com/bytedance/sonic/ast"
	"github.com/stretchr/testify/assert"
	"testing"
)

var jsonBs = []byte(`
{
    "presetBatch": {
        "assault": 120, // wtf
        "bossBully": 1
    }
}
`)

func TestJsonFileHandler(t *testing.T) {
	n, err := JsonFileHandler(jsonBs)
	assert.NoError(t, err, "handle fail")
	node := n.(ast.Node)
	fmt.Println(node.Raw())

	fmt.Println(node.Get("abc").Valid())
	fmt.Println(node.Get("presetBatch").Valid())
	fmt.Println(node.Get("abc").Exists())
	fmt.Println(node.Get("presetBatch").Exists())
}
