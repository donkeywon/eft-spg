package util

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigBox(t *testing.T) {
	d, err := ReadConfigBox()
	assert.NoError(t, err, "read fail")

	fmt.Println(jsonparser.GetString(d, "bot", "presetBatch", "assault"))
}
