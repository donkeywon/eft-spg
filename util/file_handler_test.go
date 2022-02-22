package util

import (
	"fmt"
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
	fmt.Println(GenerateSessID())
}
