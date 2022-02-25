package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerBundleRouter() {
	s.RegisterRouter("/singleplayer/bundles", s.GetBundles)
	s.RegisterRouter(".bundle", s.GetBundles)
}

func (s *Svc) GetBundles(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
