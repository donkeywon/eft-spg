package httpd

import (
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

func (s *Svc) ListFriend(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) ListChatServer(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) ListMailDialog(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetMailDialogView(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetMailDialogInfo(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) RemoveMailDialog(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) PinMailDialog(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) UnpinMailDialog(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) SetMailDialogRead(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetMailDialogAllAttachments(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) ListFriendOutbox(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) ListFriendInbox(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) SendFriendRequest(resp http.ResponseWriter, req *http.Request) {

}
