package eft

import (
	"eft-spg/util"
	"github.com/bytedance/sonic"
	"github.com/pkg/errors"
)

func Login(info []byte) (string, error) {
	n, err := sonic.Get(info)
	if err != nil {
		return "", errors.Wrap(err, util.ErrInvalidRequest)
	}

	// TODO
}
