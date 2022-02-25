package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerItemEventRouter() {
	s.RegisterRouter("/client/game/profile/items/moving", s.HandleEvents)
}

func (s *Svc) HandleEvents(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
