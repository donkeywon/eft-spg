package hook

import (
	"github.com/bytedance/sonic/ast"
)

func init() {
	RegisterPostLoadHook(insurancePostLoadProfileHook)
}

func insurancePostLoadProfileHook(p *ast.Node) {
	if p == nil {
		return
	}

	if !p.Get("insurance").Exists() {
		p.Set("insurance", ast.NewArray([]ast.Node{}))
	}
}
