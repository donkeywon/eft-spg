package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerInsuranceRouter() {
	s.RegisterRouter("/client/insurance/items/list/cost", s.ListInsuranceCost, true)
}

func (s *Svc) ListInsuranceCost(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
