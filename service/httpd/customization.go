package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerCustomizationRouter() {
	s.RegisterRouter("/client/trading/customization/storage", s.GetSuits)
	s.RegisterRouter("/client/trading/customization/{traderID}/{aaa}", s.GetTraderSuits)
}

func (s *Svc) GetSuits(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetTraderSuits(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
