package hook

import (
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"go.uber.org/multierr"
)

var _hooks []HookFn

func RegisterHook(hk HookFn) {
	_hooks = append(_hooks, hk)
}

type HookFn func(node *jsonvalue.V) error

func (hkf HookFn) DoHook(node *jsonvalue.V) error {
	return hkf(node)
}

func PostLoadHook(n *jsonvalue.V) error {
	var err error
	for _, _hook := range _hooks {
		err = multierr.Append(err, _hook.DoHook(n))
	}
	return err
}
