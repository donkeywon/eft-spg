package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerTraderRouter() {
	s.RegisterRouter("/client/trading/api/getUserAssortPrice/trader/", s.GetProfilePurchases, false)
	s.RegisterRouter("/client/trading/api/getTrader/", s.GetTrader, false)
	s.RegisterRouter("/client/trading/api/getTraderAssort/", s.GetAssort, false)
}

func (s *Svc) GetProfilePurchases(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetTrader(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetAssort(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
