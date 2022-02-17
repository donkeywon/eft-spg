package util

import (
	"fmt"
	"github.com/donkeywon/eft-spg/service/cfg/hook"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigBox(t *testing.T) {
	v, err := ReadConfigBox()
	assert.NoError(t, err, "read fail")
	err = hook.PostLoadHook(v)
	assert.NoError(t, err, "hook fail")
	fmt.Println(v.String())
}
