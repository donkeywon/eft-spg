package hook

import (
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
)

func init() {
	RegisterPostLoadHook(inraidPostLoadProfileHook)
}

func inraidPostLoadProfileHook(p *ast.Node) {
	if p == nil {
		return
	}

	if !p.Get("inraid").Exists() {
		bs := []byte(`{"location":"none","character":"none"}`)
		n, _ := sonic.Get(bs)
		p.Set("inraid", n)
	}
}
