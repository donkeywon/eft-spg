package httpd

import (
	"eft-spg/service/eft"
	"eft-spg/util"
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerCustomizationRouter() {
	s.RegisterRouter("/client/trading/customization/storage", s.GetSuits)
	s.RegisterRouter("/client/trading/customization/{traderID}/{unknown}", s.GetTraderSuits)
}

func (s *Svc) GetSuits(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return util.GetResponseWrapperFromData(eft.GetSuits(sessID)), nil
}

func (s *Svc) GetTraderSuits(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return util.GetResponseWrapperFromData(eft.GetTraderSuits(sessID, vars["traderID"])), nil
}
