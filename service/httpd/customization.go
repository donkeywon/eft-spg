package httpd

import (
	"net/http"
)

func (s *Svc) registerCustomizationRouter() {
	s.RegisterRouter("/client/trading/customization/storage", s.GetSuits)
	s.RegisterRouter("/client/trading/customization/{traderID}/{aaa}", s.GetTraderSuits)
}

func (s *Svc) GetSuits(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetTraderSuits(resp http.ResponseWriter, req *http.Request) {

}
