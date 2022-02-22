package eft

import (
	"eft-spg/service/profile"
	"eft-spg/util"
	"github.com/bytedance/sonic"
	"github.com/pkg/errors"
)

func (s *Svc) Login(info []byte) (string, error) {
	n, err := sonic.Get(info)
	if err != nil {
		return "", errors.Wrap(err, util.ErrLoginFail)
	}

	un, err := n.Get("username").String()
	if err != nil {
		return "", errors.Wrapf(err, util.ErrLoginFail)
	}

	sessID, err := profile.GetSvc().GetSessIDByUsername(un)
	if err != nil {
		return "", errors.Wrapf(err, util.ErrLoginFail)
	}

	return sessID, nil
}
