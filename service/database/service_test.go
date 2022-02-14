package database

import (
	"eft-spg/util"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Open(t *testing.T) {
	m, err := util.ReadJsonDir("/Users/donkeywon/code/other/EFT/Server/project/assets/database/")
	assert.NoError(t, err, "read dir fail")
	fmt.Println(m.Raw())
}
