package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerQuestRouter() {
	s.RegisterRouter("/client/quest/list", s.ListQuests)
	s.RegisterRouter("/client/repeatalbeQuests/activityPeriods", s.ActivityPeriods)
}

func (s *Svc) ListQuests(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil
}

func (s *Svc) ActivityPeriods(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil
}
