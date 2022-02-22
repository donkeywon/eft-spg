package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerHealthRouter() {
	s.RegisterRouter("/player/health/sync", s.SyncHealth, true)
}

func (s *Svc) SyncHealth(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
