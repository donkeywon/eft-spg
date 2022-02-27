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
	s.RegisterRouter("/launcher/server/connect", s.Connect)
	s.RegisterRouter("/launcher/profile/login", s.Login)
	s.RegisterRouter("/launcher/profile/register", s.Register)
	s.RegisterRouter("/launcher/profile/get", s.GetProfile)
	s.RegisterRouter("/launcher/profile/change/username", s.ChangeUsername)
	s.RegisterRouter("/launcher/profile/change/password", s.ChangePassword)
	s.RegisterRouter("/launcher/profile/change/wipe", s.Wipe)
	s.RegisterRouter("/launcher/profile/info", s.GetMiniProfile)
	s.RegisterRouter("/launcher/profiles", s.GetAllMiniProfiles)
	s.RegisterRouter("/launcher/server/version", s.GetServerVersion)
	s.RegisterRouter("/launcher/profile/remove", s.RemoveProfile)
	s.RegisterRouter("/launcher/profile/compatibleTarkovVersion", s.GetCompatibleTarkovVersion)
	s.RegisterRouter("/launcher/ping", s.Ping)
}

func (s *Svc) Connect(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	pe, err := database.GetProfileEditions()
	if err != nil {
		return nil, err
	}

	editions := "[\"" + strings.Join(pe, `","`) + "\"]"

	resp := fmt.Sprintf(`{"backendUrl":"%s","name":"%s","editions":%s}`, s.backendUrl(), util2.ServerName, editions)
	return util.String2Bytes(resp), nil
}

func (s *Svc) Login(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	resp := "FAILED"

	username, err := body.Get("username").String()
	if err != nil {
		return resp, errors.New(util2.ErrIllegalArg)
	}

	loginSessID, err := eft.Login(username)
	if err != nil {
		s.Error("User login fail", zap.Error(err))
	} else {
		resp = loginSessID
	}

	return resp, nil
}

func (s *Svc) Register(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	_, err := eft.Register(body)
	if err != nil {
		return "FAILED", errors.Wrap(err, util2.ErrRegisterFail)
	}
	return "OK", nil
}

func (s *Svc) GetProfile(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	username, err := body.Get("username").String()
	if err != nil {
		return "", errors.New(util2.ErrIllegalArg)
	}

	_, p := profile.GetSessProfileByUsername(username)
	if p == nil {
		return "", errors.New(util2.ErrUserNotExist)
	}

	return p.Get("info").MarshalJSON()
}

func (s *Svc) ChangeUsername(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	old, err := body.Get("username").String()
	if err != nil {
		return "FAILED", errors.New(util2.ErrIllegalArg)
	}
	newUn, err := body.Get("change").String()
	if err != nil {
		return "FAILED", errors.New(util2.ErrIllegalArg)
	}

	err = eft.ChangeUsername(old, newUn)
	if err != nil {
		return "FAILED", errors.Wrap(err, util2.ErrChangeUsername)
	}

	return "OK", nil
}

func (s *Svc) ChangePassword(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	newPw, err := body.Get("change").String()
	if err != nil {
		return "FAILED", errors.New(util2.ErrIllegalArg)
	}
	username, err := body.Get("username").String()
	if err != nil {
		return "FAILED", errors.New(util2.ErrIllegalArg)
	}

	err = eft.ChangePassword(username, newPw)
	if err != nil {
		return "FAILED", errors.Wrap(err, util2.ErrChangeUsername)
	}

	return "OK", nil
}

func (s *Svc) Wipe(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	username, err := body.Get("username").String()
	if err != nil {
		return "FAILED", errors.New(util2.ErrIllegalArg)
	}
	edition, err := body.Get("edition").String()
	if err != nil {
		return "FAILED", errors.New(util2.ErrIllegalArg)
	}

	err = eft.Wipe(username, edition)
	if err != nil {
		return "FAILED", err
	}

	return "OK", nil
}

func (s *Svc) GetMiniProfile(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return profile.GetMiniProfile(sessID)
}

func (s *Svc) GetAllMiniProfiles(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return profile.GetAllMiniProfiles()
}

func (s *Svc) Ping(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return "pong!", nil
}

func (s *Svc) GetServerVersion(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return util2.ServerVersion, nil
}

func (s *Svc) RemoveProfile(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	err := profile.RemoveProfile(sessID)
	if err != nil {
		return "false", err
	}

	return "true", nil
}

func (s *Svc) GetCompatibleTarkovVersion(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return cfg.GetCfg().GetByPath("aki", "compatibleTarkovVersion").String()
}
