package hook

import "github.com/bytedance/sonic/ast"

var _hooks []HookFn

func RegisterHook(hk HookFn) {
	_hooks = append(_hooks, hk)
}

type HookFn func(node *ast.Node) error

func (hkf HookFn) DoHook(node *ast.Node) error {
	return hkf(node)
}
