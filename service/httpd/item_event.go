package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerItemEventRouter() {
	s.RegisterRouter("/client/game/profile/items/moving", s.HandleEvents, true)
}

func (s *Svc) HandleEvents(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
