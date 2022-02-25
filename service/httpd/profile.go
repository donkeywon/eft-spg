package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerProfileRouter() {
	s.RegisterRouter("/client/game/profile/create", s.CreateProfile)
	s.RegisterRouter("/client/game/profile/list", s.GetProfileData)
	s.RegisterRouter("/client/game/profile/savage/regenerate", s.ReGenerateScav)
	s.RegisterRouter("/client/game/profile/voice/change", s.ChangeVoice)
	s.RegisterRouter("/client/game/profile/nickname/change", s.ChangeNickname)
	s.RegisterRouter("/client/game/profile/nickname/validate", s.ValidateNickname)
	s.RegisterRouter("/client/game/profile/nickname/reserved", s.GetReservedNickname)
	s.RegisterRouter("/client/profile/status", s.GetProfileStatus)
}

func (s *Svc) CreateProfile(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetProfileData(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ReGenerateScav(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ChangeVoice(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ChangeNickname(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ValidateNickname(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetReservedNickname(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetProfileStatus(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
