package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerImageRouter() {
	s.RegisterRouter(".jpg", s.GetImage)
	s.RegisterRouter(".png", s.GetImage)
	s.RegisterRouter(".ico", s.GetImage)
}

func (s *Svc) GetImage(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
