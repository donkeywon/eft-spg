package httpd

import (
	"net/http"
)

func (s *Svc) registerRagfairRouter() {
	s.RegisterRouter("/client/ragfair/search", s.Search)
	s.RegisterRouter("/client/ragfair/find", s.Search)
	s.RegisterRouter("/client/ragfair/itemMarketPrice", s.GetMarketPrice)
	s.RegisterRouter("/client/items/prices", s.GetItemPrices)
	s.RegisterRouter("/client/trading/api/traderSettings", s.GetTraderSettings)
}

func (s *Svc) Search(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetMarketPrice(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetItemPrices(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetTraderSettings(resp http.ResponseWriter, req *http.Request) {

}
