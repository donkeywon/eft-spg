package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerInraidRouter() {
	s.RegisterRouter("/raid/profile/save", s.SaveProfile)
	s.RegisterRouter("/singleplayer/settings/raid/endstate", s.GetRaidEndState)
	s.RegisterRouter("/singleplayer/settings/weapon/durability", s.GetWeaponDurability)
	s.RegisterRouter("/singleplayer/settings/raid/menu", s.GetRaidMenuSettings)
}

func (s *Svc) SaveProfile(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetRaidEndState(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetWeaponDurability(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetRaidMenuSettings(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) RegisterPlayer(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
