package config

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBotConfig(t *testing.T) {
	n, err := sonic.Get(c)
	assert.NoError(t, err, "parse config fail")
	fmt.Println(n.Raw())
}
