package httpd

import (
	"net/http"
)

func (s *Svc) registerBundleRouter() {
	s.RegisterRouter("/singleplayer/bundles", s.GetBundles)
	s.RegisterRouter(".bundle", s.GetBundles)
}

func (s *Svc) GetBundles(resp http.ResponseWriter, req *http.Request) {

}
