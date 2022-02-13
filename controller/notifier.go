package controller

import "net/http"

func init() {
	RegisterRouter("/client/notifier/channel/create", CreateNotifierChannel)
	RegisterRouter("/client/game/profile/select", SelectProfile)
	RegisterRouter("/?last_id", Notify)
	RegisterRouter("/notifierServer", Notify)
	RegisterRouter("/push/notifier/get/", GetNotifier)
}

func CreateNotifierChannel(resp http.ResponseWriter, req *http.Request) {

}

func SelectProfile(resp http.ResponseWriter, req *http.Request) {

}

func Notify(resp http.ResponseWriter, req *http.Request) {

}

func GetNotifier(resp http.ResponseWriter, req *http.Request) {

}
