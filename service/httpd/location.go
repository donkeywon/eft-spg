package httpd

import (
	"net/http"
)

func (s *Svc) registerLocationRouter() {
	s.RegisterRouter("/client/locations", s.GetLocationData)
	s.RegisterRouter("/client/location/getLocalloot", s.GetLocation)
}

func (s *Svc) GetLocationData(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetLocation(resp http.ResponseWriter, req *http.Request) {

}
