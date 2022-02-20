package httpd

import (
	"net/http"
)

func (s *Svc) registerPresetBuildRouter() {
	s.RegisterRouter("/client/handbook/builds/my/list", s.GetHandbookUserList)
}

func (s *Svc) GetHandbookUserList(resp http.ResponseWriter, req *http.Request) {

}
