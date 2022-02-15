package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigBox(t *testing.T) {
	n, err := ReadConfigBox()
	assert.NoError(t, err, "read json box fail")
	s, err := n.Raw()
	assert.NoError(t, err, "read node fail")
	fmt.Println(s)
}
