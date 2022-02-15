package hook

import (
	"eft-spg/util"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBotHook(t *testing.T) {
	c, err := util.ReadConfigBox()
	assert.NoError(t, err, "read config box fail")

	err = PostLoadHook(c)
	fmt.Println(c.Raw())
	assert.NoError(t, err, "hook config fail")
}
