package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerNotifierRouter() {
	s.RegisterRouter("/client/notifier/channel/create", s.CreateNotifierChannel)
	s.RegisterRouter("/client/game/profile/select", s.SelectProfile)
	s.RegisterRouter("/?last_id", s.Notify)
	s.RegisterRouter("/notifierServer", s.Notify)
	s.RegisterRouter("/push/notifier/get/", s.GetNotifier)
	s.RegisterRouter("/push/notifier/getwebsocket/", s.GetNotifier)
}

func (s *Svc) CreateNotifierChannel(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) SelectProfile(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) Notify(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetNotifier(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
