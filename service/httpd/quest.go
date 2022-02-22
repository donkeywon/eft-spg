package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerQuestRouter() {
	s.RegisterRouter("/client/quest/list", s.ListQuests, true)
	s.RegisterRouter("/client/repeatalbeQuests/activityPeriods", s.ActivityPeriods, true)
}

func (s *Svc) ListQuests(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil
}

func (s *Svc) ActivityPeriods(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil
}
