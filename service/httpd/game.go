package httpd

import (
	"github.com/bytedance/sonic/ast"
	"github.com/donkeywon/eft-spg/util"
	"net/http"
)

func (s *Svc) registerGameRouter() {
	s.RegisterRouter("/client/game/config", s.GetGameConfig)
	s.RegisterRouter("/client/server/list", s.ListServer)
	s.RegisterRouter("/client/game/version/validate", s.ValidateVersion)
	s.RegisterRouter("/client/game/start", s.StartGame)
	s.RegisterRouter("/client/game/logout", s.LogoutGame)
	s.RegisterRouter("/client/checkVersion", s.CheckVersion)
	s.RegisterRouter("/client/game/keepalive", s.KeepAliveGame)
	s.RegisterRouter("/singleplayer/settings/version", s.GetVersion)
}

func (s *Svc) GetGameConfig(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ListServer(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ValidateVersion(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return util.GetResponseWrapperFromData(nil), nil
}

func (s *Svc) StartGame(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) LogoutGame(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) CheckVersion(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) KeepAliveGame(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetVersion(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
