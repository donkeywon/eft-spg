package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerBundleRouter() {
	s.RegisterRouter("/singleplayer/bundles", s.GetBundles, true)
	s.RegisterRouter(".bundle", s.GetBundles, false)
}

func (s *Svc) GetBundles(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
