package httpd

import (
	"net/http"
)

func (s *Svc) registerInsuranceRouter() {
	s.RegisterRouter("/client/insurance/items/list/cost", s.ListInsuranceCost)
}

func (s *Svc) ListInsuranceCost(resp http.ResponseWriter, req *http.Request) {

}
