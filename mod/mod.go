package mod

import (
	"github.com/bytedance/sonic/ast"
	"github.com/donkeywon/eft-spg/service/database"
)

func GetValue() string {
	s, _ := database.GetDatabase().GetByPath("globals", "server", "ip").String()
	return s
}

func SetValue() {
	database.GetDatabase().GetByPath("globals", "server").Set("ip", ast.NewString("1.2.3.4"))
}
