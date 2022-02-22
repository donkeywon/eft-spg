package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerGameRouter() {
	s.RegisterRouter("/client/game/config", s.GetGameConfig, true)
	s.RegisterRouter("/client/server/list", s.ListServer, true)
	s.RegisterRouter("/client/game/version/validate", s.ValidateVersion, true)
	s.RegisterRouter("/client/game/start", s.StartGame, true)
	s.RegisterRouter("/client/game/logout", s.LogoutGame, true)
	s.RegisterRouter("/client/checkVersion", s.CheckVersion, true)
	s.RegisterRouter("/client/game/keepalive", s.KeepAliveGame, true)
	s.RegisterRouter("/singleplayer/settings/version", s.GetVersion, true)
}

func (s *Svc) GetGameConfig(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ListServer(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ValidateVersion(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) StartGame(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) LogoutGame(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) CheckVersion(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) KeepAliveGame(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetVersion(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
