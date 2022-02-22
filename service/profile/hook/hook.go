package hook

import "github.com/bytedance/sonic/ast"

var _postLoadHooks []PostLoadProfileHook

func RegisterPostLoadHook(hk PostLoadProfileHook) {
	_postLoadHooks = append(_postLoadHooks, hk)
}

type PostLoadProfileHook func(node *ast.Node)

func (h PostLoadProfileHook) DoHook(profile *ast.Node) {
	h(profile)
}

func PostLoadHook(profile *ast.Node) {
	for _, _hook := range _postLoadHooks {
		_hook.DoHook(profile)
	}
}
