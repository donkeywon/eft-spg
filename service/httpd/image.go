package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerImageRouter() {
	s.RegisterRouter(".jpg", s.GetImage, true)
	s.RegisterRouter(".png", s.GetImage, true)
	s.RegisterRouter(".ico", s.GetImage, true)
}

func (s *Svc) GetImage(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
