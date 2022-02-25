package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerLocationRouter() {
	s.RegisterRouter("/client/locations", s.GetLocationData)
	s.RegisterRouter("/client/location/getLocalloot", s.GetLocation)
}

func (s *Svc) GetLocationData(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetLocation(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
