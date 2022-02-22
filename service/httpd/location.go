package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerLocationRouter() {
	s.RegisterRouter("/client/locations", s.GetLocationData, true)
	s.RegisterRouter("/client/location/getLocalloot", s.GetLocation, false)
}

func (s *Svc) GetLocationData(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetLocation(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
