package httpd

import (
	"net/http"
)

func (s *Svc) registerNotifierRouter() {
	s.RegisterRouter("/client/notifier/channel/create", s.CreateNotifierChannel)
	s.RegisterRouter("/client/game/profile/select", s.SelectProfile)
	s.RegisterRouter("/?last_id", s.Notify)
	s.RegisterRouter("/notifierServer", s.Notify)
	s.RegisterRouter("/push/notifier/get/", s.GetNotifier)
}

func (s *Svc) CreateNotifierChannel(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) SelectProfile(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) Notify(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetNotifier(resp http.ResponseWriter, req *http.Request) {

}
