package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerRagfairRouter() {
	s.RegisterRouter("/client/ragfair/search", s.Search, true)
	s.RegisterRouter("/client/ragfair/find", s.Search, true)
	s.RegisterRouter("/client/ragfair/itemMarketPrice", s.GetMarketPrice, true)
	s.RegisterRouter("/client/items/prices", s.GetItemPrices, true)
	s.RegisterRouter("/client/trading/api/traderSettings", s.GetTraderSettings, true)
}

func (s *Svc) Search(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetMarketPrice(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetItemPrices(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetTraderSettings(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
