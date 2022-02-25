package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerInsuranceRouter() {
	s.RegisterRouter("/client/insurance/items/list/cost", s.ListInsuranceCost)
}

func (s *Svc) ListInsuranceCost(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
