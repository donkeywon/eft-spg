package httpd

import (
	"net/http"
)

func (s *Svc) registerTraderRouter() {
	s.RegisterRouter("/client/trading/api/getUserAssortPrice/trader/{traderID}", s.GetProfilePurchases)
	s.RegisterRouter("/client/trading/api/getTrader/{traderID}", s.GetTrader)
	s.RegisterRouter("/client/trading/api/getTraderAssort/{traderID}", s.GetAssort)
}

func (s *Svc) GetProfilePurchases(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetTrader(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetAssort(resp http.ResponseWriter, req *http.Request) {

}
