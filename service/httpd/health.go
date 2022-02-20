package httpd

import (
	"net/http"
)

func (s *Svc) registerHealthRouter() {
	s.RegisterRouter("/player/health/sync", s.SyncHealth)
}

func (s *Svc) SyncHealth(resp http.ResponseWriter, req *http.Request) {

}
