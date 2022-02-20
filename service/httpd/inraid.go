package httpd

import (
	"net/http"
)

func (s *Svc) registerInraidRouter() {
	s.RegisterRouter("/raid/profile/save", s.SaveProfile)
	s.RegisterRouter("/singleplayer/settings/raid/endstate", s.GetRaidEndState)
	s.RegisterRouter("/singleplayer/settings/weapon/durability", s.GetWeaponDurability)
	s.RegisterRouter("/singleplayer/settings/raid/menu", s.GetRaidMenuSettings)
}

func (s *Svc) SaveProfile(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetRaidEndState(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetWeaponDurability(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetRaidMenuSettings(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) RegisterPlayer(resp http.ResponseWriter, req *http.Request) {

}
