package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerDialogRouter() {
	s.RegisterRouter("/client/friend/list", s.ListFriend)
	s.RegisterRouter("/client/chatServer/list", s.ListChatServer)
	s.RegisterRouter("/client/mail/dialog/list", s.ListMailDialog)
	s.RegisterRouter("/client/mail/dialog/view", s.GetMailDialogView)
	s.RegisterRouter("/client/mail/dialog/info", s.GetMailDialogInfo)
	s.RegisterRouter("/client/mail/dialog/remove", s.RemoveMailDialog)
	s.RegisterRouter("/client/mail/dialog/pin", s.PinMailDialog)
	s.RegisterRouter("/client/mail/dialog/unpin", s.UnpinMailDialog)
	s.RegisterRouter("/client/mail/dialog/read", s.SetMailDialogRead)
	s.RegisterRouter("/client/mail/dialog/getAllAttachments", s.GetMailDialogAllAttachments)
	s.RegisterRouter("/client/friend/request/list/outbox", s.ListFriendOutbox)
	s.RegisterRouter("/client/friend/request/list/inbox", s.ListFriendInbox)
	s.RegisterRouter("/client/friend/request/send", s.SendFriendRequest)
}

func (s *Svc) ListFriend(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ListChatServer(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ListMailDialog(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetMailDialogView(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetMailDialogInfo(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) RemoveMailDialog(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) PinMailDialog(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) UnpinMailDialog(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) SetMailDialogRead(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetMailDialogAllAttachments(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ListFriendOutbox(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ListFriendInbox(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) SendFriendRequest(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
