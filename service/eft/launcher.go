package eft

import (
	"eft-spg/service/profile"
	"eft-spg/util"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	util2 "github.com/donkeywon/gtil/util"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
)

type userInfo struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Wipe     bool   `json:"wipe"`
	Edition  string `json:"edition"`
}

func (s *Svc) Login(username string) (string, error) {
	sessID, _ := profile.GetSvc().GetSessProfileByUsername(username)
	if sessID == "" {
		return "", errors.New(util.ErrUserNotExist)
	}

	return sessID, nil
}

func (s *Svc) Register(info *ast.Node) (string, error) {
	var username, password, edition string
	username, err := info.Get("username").String()
	password, err1 := info.Get("password").String()
	edition, err2 := info.Get("edition").String()

	if multierr.Combine(err, err1, err2) != nil {
		return "", errors.Wrap(err, util.ErrIllegalArg)
	}

	sessID, _ := profile.GetSvc().GetSessProfileByUsername(username)
	if sessID != "" {
		return "", errors.New(util.ErrUserExist)
	}

	s.createAccount(username, password, edition)
	return "", nil
}

func (s *Svc) createAccount(username string, password string, edition string) {
	sessID := util.GenerateSessID()

	info := fmt.Sprintf(`{
"info": {
	"id": "%s",
	"username": "%s",
	"password": "%s",
	"wipe": true,
	"edition": "%s"
}
}`, sessID, username, password, edition)
	p, _ := sonic.Get(util2.String2Bytes(info))

	profile.GetSvc().SetProfile(sessID, &p)
}
