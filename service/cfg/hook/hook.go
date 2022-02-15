package hook

import (
	"go.uber.org/multierr"
)

var _hooks []HookFn

func RegisterHook(hk HookFn) {
	_hooks = append(_hooks, hk)
}

type HookFn func(node []byte) error

func (hkf HookFn) DoHook(node []byte) error {
	return hkf(node)
}

func PostLoadHook(n []byte) error {
	var err error
	for _, _hook := range _hooks {
		err = multierr.Append(err, _hook.DoHook(n))
	}
	return err
}
