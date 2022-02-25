package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerHealthRouter() {
	s.RegisterRouter("/player/health/sync", s.SyncHealth)
}

func (s *Svc) SyncHealth(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
