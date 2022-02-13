package controller

import "net/http"

func init() {
	RegisterRouter("/client/game/config", GetGameConfig)
	RegisterRouter("/client/server/list", ListServer)
	RegisterRouter("/client/game/version/validate", ValidateVersion)
	RegisterRouter("/client/game/start", StartGame)
	RegisterRouter("/client/game/logout", LogoutGame)
	RegisterRouter("/client/checkVersion", CheckVersion)
	RegisterRouter("/client/game/keepalive", KeepAliveGame)
	RegisterRouter("/singleplayer/settings/version", GetVersion)
}

func GetGameConfig(resp http.ResponseWriter, req *http.Request) {

}

func ListServer(resp http.ResponseWriter, req *http.Request) {

}

func ValidateVersion(resp http.ResponseWriter, req *http.Request) {

}

func StartGame(resp http.ResponseWriter, req *http.Request) {

}

func LogoutGame(resp http.ResponseWriter, req *http.Request) {

}

func CheckVersion(resp http.ResponseWriter, req *http.Request) {

}

func KeepAliveGame(resp http.ResponseWriter, req *http.Request) {

}

func GetVersion(resp http.ResponseWriter, req *http.Request) {

}
