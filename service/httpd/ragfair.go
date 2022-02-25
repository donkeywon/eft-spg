package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerRagfairRouter() {
	s.RegisterRouter("/client/ragfair/search", s.Search)
	s.RegisterRouter("/client/ragfair/find", s.Search)
	s.RegisterRouter("/client/ragfair/itemMarketPrice", s.GetMarketPrice)
	s.RegisterRouter("/client/items/prices", s.GetItemPrices)
	s.RegisterRouter("/client/trading/api/traderSettings", s.GetTraderSettings)
}

func (s *Svc) Search(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetMarketPrice(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetItemPrices(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetTraderSettings(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
