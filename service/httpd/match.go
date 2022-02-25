package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerMatchRouter() {
	s.RegisterRouter("/raid/profile/list", s.ListRaidProfile)
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

func (s *Svc) ListRaidProfile(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ServerAvailable(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) UpdatePing(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) JoinMatch(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ExitMatch(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) CreateGroup(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) DeleteGroup(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetGroupStatus(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ExitToMenu(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) StartGroupSearch(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) StopGroupSearch(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) SendGroupInvite(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) AcceptGroupInvite(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) CancelGroupInvite(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) StartOfflineRaid(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) EndOfflineRaid(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) PutMetrics(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetMetrics(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
