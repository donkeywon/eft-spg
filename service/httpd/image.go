package httpd

import (
	"net/http"
)

func (s *Svc) registerImageRouter() {
	s.RegisterRouter(".jpg", s.GetImage)
	s.RegisterRouter(".png", s.GetImage)
	s.RegisterRouter(".ico", s.GetImage)
}

func (s *Svc) GetImage(resp http.ResponseWriter, req *http.Request) {

}
