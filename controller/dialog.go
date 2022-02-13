package controller

import "net/http"

func init() {
	RegisterRouter("/client/friend/list", ListFriend)
	RegisterRouter("/client/chatServer/list", ListChatServer)
	RegisterRouter("/client/mail/dialog/list", ListMailDialog)
	RegisterRouter("/client/mail/dialog/view", GetMailDialogView)
	RegisterRouter("/client/mail/dialog/info", GetMailDialogInfo)
	RegisterRouter("/client/mail/dialog/remove", RemoveMailDialog)
	RegisterRouter("/client/mail/dialog/pin", PinMailDialog)
	RegisterRouter("/client/mail/dialog/unpin", UnpinMailDialog)
	RegisterRouter("/client/mail/dialog/read", SetMailDialogRead)
	RegisterRouter("/client/mail/dialog/getAllAttachments", GetMailDialogAllAttachments)
	RegisterRouter("/client/friend/request/list/outbox", ListFriendOutbox)
	RegisterRouter("/client/friend/request/list/inbox", ListFriendInbox)
	RegisterRouter("/client/friend/request/send", SendFriendRequest)

}

func ListFriend(resp http.ResponseWriter, req *http.Request) {

}

func ListChatServer(resp http.ResponseWriter, req *http.Request) {

}

func ListMailDialog(resp http.ResponseWriter, req *http.Request) {

}

func GetMailDialogView(resp http.ResponseWriter, req *http.Request) {

}

func GetMailDialogInfo(resp http.ResponseWriter, req *http.Request) {

}

func RemoveMailDialog(resp http.ResponseWriter, req *http.Request) {

}

func PinMailDialog(resp http.ResponseWriter, req *http.Request) {

}

func UnpinMailDialog(resp http.ResponseWriter, req *http.Request) {

}

func SetMailDialogRead(resp http.ResponseWriter, req *http.Request) {

}

func GetMailDialogAllAttachments(resp http.ResponseWriter, req *http.Request) {

}

func ListFriendOutbox(resp http.ResponseWriter, req *http.Request) {

}

func ListFriendInbox(resp http.ResponseWriter, req *http.Request) {

}

func SendFriendRequest(resp http.ResponseWriter, req *http.Request) {

}
