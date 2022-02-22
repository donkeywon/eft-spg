package httpd

import (
	"eft-spg/service/database"
	"eft-spg/service/eft"
	"fmt"
	"github.com/bytedance/sonic/ast"
	"github.com/donkeywon/gtil/util"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func (s *Svc) registerLauncherRouter() {
	s.RegisterRouter("/launcher/server/connect", s.Connect, true)
	s.RegisterRouter("/launcher/profile/login", s.Login, true)
	s.RegisterRouter("/launcher/profile/register", s.Register, true)
	s.RegisterRouter("/launcher/profile/get", s.Get, true)
	s.RegisterRouter("/launcher/profile/change/username", s.ChangeUsername, true)
	s.RegisterRouter("/launcher/profile/change/password", s.ChangePassword, true)
	s.RegisterRouter("/launcher/profile/change/wipe", s.Wipe, true)
	s.RegisterRouter("/launcher/profile/info", s.GetMiniProfile, true)
	s.RegisterRouter("/launcher/ping", s.Ping, true)
}

func (s *Svc) Connect(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	pe, err := database.GetSvc().GetProfileEditions()
	if err != nil {
		return nil, err
	}

	editions := "[\"" + strings.Join(pe, `","`) + "\"]"

	resp := fmt.Sprintf(`{"backendUrl":"%s","name":"%s","editions":%s}`, s.backendUrl(), ServerName, editions)
	return util.String2Bytes(resp), nil
}

func (s *Svc) Login(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	resp := "FAILED"

	username, err := body.Get("username").String()
	if err != nil {
		return resp, nil
	}

	resp, err = eft.GetSvc().Login(username)
	if err != nil {
		s.Error("User login fail", zap.Error(err))
	}

	return resp, nil
}

func (s *Svc) Register(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) Get(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ChangeUsername(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ChangePassword(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) Wipe(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetMiniProfile(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) Ping(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
