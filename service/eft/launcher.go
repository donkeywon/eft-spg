package eft

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/donkeywon/eft-spg/service/database"
	"github.com/donkeywon/eft-spg/service/profile"
	"github.com/donkeywon/eft-spg/util"
	util2 "github.com/donkeywon/gtil/util"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
)

func Login(username string) (string, error) {
	sessID, _ := profile.GetSessProfileByUsername(username)
	if sessID == "" {
		return "", errors.New(util.ErrUserNotExist)
	}

	return sessID, nil
}

func Register(info *ast.Node) (string, error) {
	var username, password, edition string
	username, err := info.Get("username").String()
	password, err1 := info.Get("password").String()
	edition, err2 := info.Get("edition").String()

	if multierr.Combine(err, err1, err2) != nil {
		return "", errors.Wrap(err, util.ErrIllegalArg)
	}

	et := database.GetProfileEditionsTemplate()
	if !et.Get(edition).Exists() {
		return "", errors.New(util.ErrIllegalArg)
	}

	sessID, _ := profile.GetSessProfileByUsername(username)
	if sessID != "" {
		return "", errors.New(util.ErrUserExist)
	}

	createAccount(username, password, edition)
	return "", nil
}

func createAccount(username string, password string, edition string) {
	sessID := util.GenerateID()

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

	profile.SetProfile(sessID, &p)
	profile.LoadProfile(sessID)
	profile.SaveProfile(sessID)
}

func ChangeUsername(old string, new string) error {
	_, p := profile.GetSessProfileByUsername(old)
	if p == nil {
		return errors.New(util.ErrUserNotExist)
	}

	p.Get("info").Set("username", ast.NewString(new))
	return nil
}

func ChangePassword(username string, new string) error {
	_, p := profile.GetSessProfileByUsername(username)
	if p == nil {
		return errors.New(util.ErrUserNotExist)
	}

	p.Get("info").Set("password", ast.NewString(new))
	return nil
}

func Wipe(username string, edition string) error {
	_, p := profile.GetSessProfileByUsername(username)
	if p == nil {
		return errors.New(util.ErrUserNotExist)
	}

	et := database.GetProfileEditionsTemplate()
	if !et.Get(edition).Exists() {
		return errors.New(util.ErrIllegalArg)
	}

	p.Get("info").Set("edition", ast.NewString(edition))
	p.Get("info").Set("wipe", ast.NewBool(true))
	return nil
}
