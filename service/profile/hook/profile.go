package hook

import (
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
)

func init() {
	RegisterPostLoadHook(profilePostLoadProfileHook)
}

func profilePostLoadProfileHook(p *ast.Node) {
	if p == nil {
		return
	}

	if !p.Get("characters").Exists() {
		n, _ := sonic.Get([]byte(`{"pmc":{},"scav":{}}`))
		p.Set("characters", n)
	}
}
