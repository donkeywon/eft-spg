package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerPresetBuildRouter() {
	s.RegisterRouter("/client/handbook/builds/my/list", s.GetHandbookUserList, true)
}

func (s *Svc) GetHandbookUserList(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
