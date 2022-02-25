package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerTraderRouter() {
	s.RegisterRouter("/client/trading/api/getUserAssortPrice/trader/{traderID}", s.GetProfilePurchases)
	s.RegisterRouter("/client/trading/api/getTrader/{traderID}", s.GetTrader)
	s.RegisterRouter("/client/trading/api/getTraderAssort/{traderID}", s.GetAssort)
}

func (s *Svc) GetProfilePurchases(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetTrader(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetAssort(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
