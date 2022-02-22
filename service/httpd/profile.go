package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerProfileRouter() {
	s.RegisterRouter("/client/game/profile/create", s.CreateProfile, true)
	s.RegisterRouter("/client/game/profile/list", s.GetProfileData, true)
	s.RegisterRouter("/client/game/profile/savage/regenerate", s.ReGenerateScav, true)
	s.RegisterRouter("/client/game/profile/voice/change", s.ChangeVoice, true)
	s.RegisterRouter("/client/game/profile/nickname/change", s.ChangeNickname, true)
	s.RegisterRouter("/client/game/profile/nickname/validate", s.ValidateNickname, true)
	s.RegisterRouter("/client/game/profile/nickname/reserved", s.GetReservedNickname, true)
	s.RegisterRouter("/client/profile/status", s.GetProfileStatus, true)
}

func (s *Svc) CreateProfile(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetProfileData(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ReGenerateScav(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ChangeVoice(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ChangeNickname(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ValidateNickname(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetReservedNickname(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetProfileStatus(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
