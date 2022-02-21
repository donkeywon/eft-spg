package hook

import (
	"github.com/bytedance/sonic/ast"
	"go.uber.org/multierr"
)

var _hooks []HookFn

func RegisterHook(hk HookFn) {
	_hooks = append(_hooks, hk)
}

type HookFn func(node *ast.Node) error

func (hkf HookFn) DoHook(node *ast.Node) error {
	return hkf(node)
}

func PostLoadHook(n *ast.Node) error {
	var err error
	for _, _hook := range _hooks {
		err = multierr.Append(err, _hook.DoHook(n))
	}
	return err
}
