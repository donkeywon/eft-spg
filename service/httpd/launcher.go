package httpd

import (
	"eft-spg/service/database"
	"eft-spg/util"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func (s *Svc) registerLauncherRouter() {
	s.RegisterRouter("/launcher/server/connect", s.Connect)
	s.RegisterRouter("/launcher/profile/login", s.Login)
	s.RegisterRouter("/launcher/profile/register", s.Register)
	s.RegisterRouter("/launcher/profile/get", s.Get)
	s.RegisterRouter("/launcher/profile/change/username", s.ChangeUsername)
	s.RegisterRouter("/launcher/profile/change/password", s.ChangePassword)
	s.RegisterRouter("/launcher/profile/change/wipe", s.Wipe)
	s.RegisterRouter("/launcher/profile/info", s.GetMiniProfile)
	s.RegisterRouter("/launcher/ping", s.Ping)
}

func (s *Svc) Connect(w http.ResponseWriter, r *http.Request) {
	pe, err := database.GetProfileEditions()
	if err != nil {
		s.Error("Get profile editions fail", zap.Error(err))
		return
	}

	editions := "[\"" + strings.Join(pe, `","`) + "\"]"

	body := fmt.Sprintf(`{"backendUrl":"%s","name":"%s","editions":%s`, s.backendUrl(), ServerName, editions)
	s.logRespErr(util.DoResponseJsonString(body, w), r)
}

func (s *Svc) Login(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) Register(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) Get(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) ChangeUsername(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) ChangePassword(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) Wipe(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetMiniProfile(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) Ping(resp http.ResponseWriter, req *http.Request) {

}
