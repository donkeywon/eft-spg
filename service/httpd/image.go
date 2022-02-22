package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerImageRouter() {
	s.RegisterRouter(".jpg", s.GetImage, false)
	s.RegisterRouter(".png", s.GetImage, false)
	s.RegisterRouter(".ico", s.GetImage, false)
}

func (s *Svc) GetImage(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
