package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerNotifierRouter() {
	s.RegisterRouter("/client/notifier/channel/create", s.CreateNotifierChannel, true)
	s.RegisterRouter("/client/game/profile/select", s.SelectProfile, true)
	s.RegisterRouter("/?last_id", s.Notify, false)
	s.RegisterRouter("/notifierServer", s.Notify, false)
	s.RegisterRouter("/push/notifier/get/", s.GetNotifier, false)
	s.RegisterRouter("/push/notifier/getwebsocket/", s.GetNotifier, false)
}

func (s *Svc) CreateNotifierChannel(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) SelectProfile(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) Notify(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetNotifier(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
