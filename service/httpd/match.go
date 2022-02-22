package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerMatchRouter() {
	s.RegisterRouter("/raid/profile/list", s.ListRaidProfile, true)
	s.RegisterRouter("/client/match/available", s.ServerAvailable, true)
	s.RegisterRouter("/client/match/updatePing", s.UpdatePing, true)
	s.RegisterRouter("/client/match/join", s.JoinMatch, true)
	s.RegisterRouter("/client/match/group/start_game", s.JoinMatch, true)
	s.RegisterRouter("/client/match/exit", s.ExitMatch, true)
	s.RegisterRouter("/client/match/group/create", s.CreateGroup, true)
	s.RegisterRouter("/client/match/group/delete", s.DeleteGroup, true)
	s.RegisterRouter("/client/match/group/status", s.GetGroupStatus, true)
	s.RegisterRouter("/client/match/group/exit_from_menu", s.ExitToMenu, true)
	s.RegisterRouter("/client/match/group/looking/start", s.StartGroupSearch, true)
	s.RegisterRouter("/client/match/group/looking/stop", s.StopGroupSearch, true)
	s.RegisterRouter("/client/match/group/invite/send", s.SendGroupInvite, true)
	s.RegisterRouter("/client/match/group/invite/accept", s.AcceptGroupInvite, true)
	s.RegisterRouter("/client/match/group/invite/cancel", s.CancelGroupInvite, true)
	s.RegisterRouter("/client/match/offline/start", s.StartOfflineRaid, true)
	s.RegisterRouter("/client/match/offline/end", s.EndOfflineRaid, true)
	s.RegisterRouter("/client/putMetrics", s.PutMetrics, true)
	s.RegisterRouter("/client/getMetricsConfig", s.GetMetrics, true)
}

func (s *Svc) ListRaidProfile(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ServerAvailable(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) UpdatePing(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) JoinMatch(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ExitMatch(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) CreateGroup(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) DeleteGroup(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetGroupStatus(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ExitToMenu(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) StartGroupSearch(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) StopGroupSearch(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) SendGroupInvite(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) AcceptGroupInvite(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) CancelGroupInvite(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) StartOfflineRaid(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) EndOfflineRaid(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) PutMetrics(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetMetrics(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
