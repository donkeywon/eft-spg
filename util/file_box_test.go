package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigBox(t *testing.T) {
	//DatabaseFileBox.Walk(func(s string, file packd.File) error {
	//	fmt.Println(s)
	//	return nil
	//})

	n, err := ReadJsonBox(DatabaseFileBox)
	assert.NoError(t, err, "read json box fail")
	s, err := n.Raw()
	assert.NoError(t, err, "read node fail")
	fmt.Println(s)
}

func clearMap(m map[string]interface{}) {
	for k, v := range m {
		if _, ok := v.(map[string]interface{}); ok {
			clearMap(m[k].(map[string]interface{}))
		} else {
			m[k] = ""
		}
	}
}
