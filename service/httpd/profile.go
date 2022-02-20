package httpd

import (
	"net/http"
)

func (s *Svc) registerProfileRouter() {
	s.RegisterRouter("/client/game/profile/create", s.CreateProfile)
	s.RegisterRouter("/client/game/profile/list", s.GetProfileData)
	s.RegisterRouter("/client/game/profile/savage/regenerate", s.ReGenerateScav)
	s.RegisterRouter("/client/game/profile/voice/change", s.ChangeVoice)
	s.RegisterRouter("/client/game/profile/nickname/change", s.ChangeNickname)
	s.RegisterRouter("/client/game/profile/nickname/validate", s.ValidateNickname)
	s.RegisterRouter("/client/game/profile/nickname/reserved", s.GetReservedNickname)
	s.RegisterRouter("/client/profile/status", s.GetProfileStatus)
}

func (s *Svc) CreateProfile(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetProfileData(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) ReGenerateScav(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) ChangeVoice(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) ChangeNickname(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) ValidateNickname(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetReservedNickname(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetProfileStatus(resp http.ResponseWriter, req *http.Request) {

}
