package httpd

import (
	"eft-spg/service/cfg"
	"eft-spg/service/database"
	"eft-spg/service/eft"
	"eft-spg/service/profile"
	util2 "eft-spg/util"
	"fmt"
	"github.com/bytedance/sonic/ast"
	"github.com/donkeywon/gtil/util"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func (s *Svc) registerLauncherRouter() {
	s.RegisterRouter("/launcher/server/connect", s.Connect, true)
	s.RegisterRouter("/launcher/profile/login", s.Login, true)
	s.RegisterRouter("/launcher/profile/register", s.Register, true)
	s.RegisterRouter("/launcher/profile/get", s.GetProfile, true)
	s.RegisterRouter("/launcher/profile/change/username", s.ChangeUsername, true)
	s.RegisterRouter("/launcher/profile/change/password", s.ChangePassword, true)
	s.RegisterRouter("/launcher/profile/change/wipe", s.Wipe, true)
	s.RegisterRouter("/launcher/profile/info", s.GetMiniProfile, true)
	s.RegisterRouter("/launcher/profiles", s.GetAllMiniProfiles, true)
	s.RegisterRouter("/launcher/server/version", s.GetServerVersion, true)
	s.RegisterRouter("/launcher/profile/remove", s.RemoveProfile, true)
	s.RegisterRouter("/launcher/profile/compatibleTarkovVersion", s.GetCompatibleTarkovVersion, true)
	s.RegisterRouter("/launcher/ping", s.Ping, true)
}

func (s *Svc) Connect(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	pe, err := database.GetSvc().GetProfileEditions()
	if err != nil {
		return nil, err
	}

	editions := "[\"" + strings.Join(pe, `","`) + "\"]"

	resp := fmt.Sprintf(`{"backendUrl":"%s","name":"%s","editions":%s}`, s.backendUrl(), eft.ServerName, editions)
	return util.String2Bytes(resp), nil
}

func (s *Svc) Login(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	resp := "FAILED"

	username, err := body.Get("username").String()
	if err != nil {
		return resp, errors.New(util2.ErrIllegalArg)
	}

	resp, err = eft.GetSvc().Login(username)
	if err != nil {
		s.Error("User login fail", zap.Error(err))
	}

	return resp, nil
}

func (s *Svc) Register(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	_, err := eft.GetSvc().Register(body)
	if err != nil {
		return "FAILED", errors.Wrap(err, util2.ErrRegisterFail)
	}
	return "OK", nil
}

func (s *Svc) GetProfile(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	username, err := body.Get("username").String()
	if err != nil {
		return "", errors.New(util2.ErrIllegalArg)
	}

	_, p := profile.GetSvc().GetSessProfileByUsername(username)
	if p == nil {
		return "", errors.New(util2.ErrUserNotExist)
	}

	return p.Get("info").MarshalJSON()
}

func (s *Svc) ChangeUsername(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	old, err := body.Get("username").String()
	if err != nil {
		return "FAILED", errors.New(util2.ErrIllegalArg)
	}
	newUn, err := body.Get("change").String()
	if err != nil {
		return "FAILED", errors.New(util2.ErrIllegalArg)
	}

	err = eft.GetSvc().ChangeUsername(old, newUn)
	if err != nil {
		return "FAILED", errors.Wrap(err, util2.ErrChangeUsername)
	}

	return "OK", nil
}

func (s *Svc) ChangePassword(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	newPw, err := body.Get("change").String()
	if err != nil {
		return "FAILED", errors.New(util2.ErrIllegalArg)
	}
	username, err := body.Get("username").String()
	if err != nil {
		return "FAILED", errors.New(util2.ErrIllegalArg)
	}

	err = eft.GetSvc().ChangePassword(username, newPw)
	if err != nil {
		return "FAILED", errors.Wrap(err, util2.ErrChangeUsername)
	}

	return "OK", nil
}

func (s *Svc) Wipe(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	username, err := body.Get("username").String()
	if err != nil {
		return "FAILED", errors.New(util2.ErrIllegalArg)
	}
	edition, err := body.Get("edition").String()
	if err != nil {
		return "FAILED", errors.New(util2.ErrIllegalArg)
	}

	err = eft.GetSvc().Wipe(username, edition)
	if err != nil {
		return "FAILED", err
	}

	return "OK", nil
}

func (s *Svc) GetMiniProfile(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return profile.GetSvc().GetMiniProfile(sessID)
}

func (s *Svc) GetAllMiniProfiles(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return profile.GetSvc().GetAllMiniProfiles()
}

func (s *Svc) Ping(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return "pong!", nil
}

func (s *Svc) GetServerVersion(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return eft.Version, nil
}

func (s *Svc) RemoveProfile(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	err := profile.GetSvc().RemoveProfile(sessID)
	if err != nil {
		return "false", err
	}

	return "true", nil
}

func (s *Svc) GetCompatibleTarkovVersion(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return cfg.GetSvc().GetConfig().GetByPath("aki", "compatibleTarkovVersion").String()
}
