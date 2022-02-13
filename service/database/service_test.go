package database

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Open(t *testing.T) {
	m, err := readDir("/Users/donkeywon/code/other/EFT/Server/project/assets/database/")
	assert.NoError(t, err, "read dir fail")
	fmt.Println(m.Raw())
}
