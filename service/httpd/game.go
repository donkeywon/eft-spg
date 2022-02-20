package httpd

import (
	"net/http"
)

func (s *Svc) registerGameRouter() {
	s.RegisterRouter("/client/game/config", s.GetGameConfig)
	s.RegisterRouter("/client/server/list", s.ListServer)
	s.RegisterRouter("/client/game/version/validate", s.ValidateVersion)
	s.RegisterRouter("/client/game/start", s.StartGame)
	s.RegisterRouter("/client/game/logout", s.LogoutGame)
	s.RegisterRouter("/client/checkVersion", s.CheckVersion)
	s.RegisterRouter("/client/game/keepalive", s.KeepAliveGame)
	s.RegisterRouter("/singleplayer/settings/version", s.GetVersion)
}

func (s *Svc) GetGameConfig(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) ListServer(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) ValidateVersion(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) StartGame(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) LogoutGame(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) CheckVersion(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) KeepAliveGame(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetVersion(resp http.ResponseWriter, req *http.Request) {

}
