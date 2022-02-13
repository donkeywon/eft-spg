package controller

import "net/http"

func init() {
	RegisterRouter("/launcher/server/connect", Connect)
	RegisterRouter("/launcher/profile/login", Login)
	RegisterRouter("/launcher/profile/register", Register)
	RegisterRouter("/launcher/profile/get", Get)
	RegisterRouter("/launcher/profile/change/username", ChangeUsername)
	RegisterRouter("/launcher/profile/change/password", ChangePassword)
	RegisterRouter("/launcher/profile/change/wipe", Wipe)
	RegisterRouter("/launcher/profile/info", GetMiniProfile)
	RegisterRouter("/launcher/ping", Ping)
}

func Connect(resp http.ResponseWriter, req *http.Request) {

}

func Login(resp http.ResponseWriter, req *http.Request) {

}

func Register(resp http.ResponseWriter, req *http.Request) {

}

func Get(resp http.ResponseWriter, req *http.Request) {

}

func ChangeUsername(resp http.ResponseWriter, req *http.Request) {

}

func ChangePassword(resp http.ResponseWriter, req *http.Request) {

}

func Wipe(resp http.ResponseWriter, req *http.Request) {

}

func GetMiniProfile(resp http.ResponseWriter, req *http.Request) {

}

func Ping(resp http.ResponseWriter, req *http.Request) {

}
