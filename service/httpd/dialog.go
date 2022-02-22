package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerDialogRouter() {
	s.RegisterRouter("/client/friend/list", s.ListFriend, true)
	s.RegisterRouter("/client/chatServer/list", s.ListChatServer, true)
	s.RegisterRouter("/client/mail/dialog/list", s.ListMailDialog, true)
	s.RegisterRouter("/client/mail/dialog/view", s.GetMailDialogView, true)
	s.RegisterRouter("/client/mail/dialog/info", s.GetMailDialogInfo, true)
	s.RegisterRouter("/client/mail/dialog/remove", s.RemoveMailDialog, true)
	s.RegisterRouter("/client/mail/dialog/pin", s.PinMailDialog, true)
	s.RegisterRouter("/client/mail/dialog/unpin", s.UnpinMailDialog, true)
	s.RegisterRouter("/client/mail/dialog/read", s.SetMailDialogRead, true)
	s.RegisterRouter("/client/mail/dialog/getAllAttachments", s.GetMailDialogAllAttachments, true)
	s.RegisterRouter("/client/friend/request/list/outbox", s.ListFriendOutbox, true)
	s.RegisterRouter("/client/friend/request/list/inbox", s.ListFriendInbox, true)
	s.RegisterRouter("/client/friend/request/send", s.SendFriendRequest, true)
}

func (s *Svc) ListFriend(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ListChatServer(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ListMailDialog(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetMailDialogView(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetMailDialogInfo(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) RemoveMailDialog(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) PinMailDialog(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) UnpinMailDialog(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) SetMailDialogRead(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetMailDialogAllAttachments(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ListFriendOutbox(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) ListFriendInbox(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) SendFriendRequest(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
