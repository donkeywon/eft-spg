package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerInraidRouter() {
	s.RegisterRouter("/raid/profile/save", s.SaveProfile, true)
	s.RegisterRouter("/singleplayer/settings/raid/endstate", s.GetRaidEndState, true)
	s.RegisterRouter("/singleplayer/settings/weapon/durability", s.GetWeaponDurability, true)
	s.RegisterRouter("/singleplayer/settings/raid/menu", s.GetRaidMenuSettings, true)
}

func (s *Svc) SaveProfile(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetRaidEndState(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetWeaponDurability(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetRaidMenuSettings(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) RegisterPlayer(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
