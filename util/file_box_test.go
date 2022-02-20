package util

import (
	"eft-spg/service/cfg/hook"
	"fmt"
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
