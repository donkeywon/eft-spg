package httpd

import (
	"net/http"
)

func (s *Svc) registerMatchRouter() {
	s.RegisterRouter("/raid/profile/list", s.GetProfile)
	s.RegisterRouter("/client/match/available", s.ServerAvailable)
	s.RegisterRouter("/client/match/updatePing", s.UpdatePing)
	s.RegisterRouter("/client/match/join", s.JoinMatch)
	s.RegisterRouter("/client/match/group/start_game", s.JoinMatch)
	s.RegisterRouter("/client/match/exit", s.ExitMatch)
	s.RegisterRouter("/client/match/group/create", s.CreateGroup)
	s.RegisterRouter("/client/match/group/delete", s.DeleteGroup)
	s.RegisterRouter("/client/match/group/status", s.GetGroupStatus)
	s.RegisterRouter("/client/match/group/exit_from_menu", s.ExitToMenu)
	s.RegisterRouter("/client/match/group/looking/start", s.StartGroupSearch)
	s.RegisterRouter("/client/match/group/looking/stop", s.StopGroupSearch)
	s.RegisterRouter("/client/match/group/invite/send", s.SendGroupInvite)
	s.RegisterRouter("/client/match/group/invite/accept", s.AcceptGroupInvite)
	s.RegisterRouter("/client/match/group/invite/cancel", s.CancelGroupInvite)
	s.RegisterRouter("/client/match/offline/start", s.StartOfflineRaid)
	s.RegisterRouter("/client/match/offline/end", s.EndOfflineRaid)
	s.RegisterRouter("/client/putMetrics", s.PutMetrics)
	s.RegisterRouter("/client/getMetricsConfig", s.GetMetrics)
}

func (s *Svc) GetProfile(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) ServerAvailable(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) UpdatePing(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) JoinMatch(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) ExitMatch(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) CreateGroup(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) DeleteGroup(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetGroupStatus(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) ExitToMenu(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) StartGroupSearch(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) StopGroupSearch(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) SendGroupInvite(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) AcceptGroupInvite(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) CancelGroupInvite(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) StartOfflineRaid(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) EndOfflineRaid(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) PutMetrics(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetMetrics(resp http.ResponseWriter, req *http.Request) {

}
