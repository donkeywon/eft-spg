package controller

import "net/http"

func init() {
	RegisterRouter("/client/game/profile/create", CreateProfile)
	RegisterRouter("/client/game/profile/list", GetProfileData)
	RegisterRouter("/client/game/profile/savage/regenerate", ReGenerateScav)
	RegisterRouter("/client/game/profile/voice/change", ChangeVoice)
	RegisterRouter("/client/game/profile/nickname/change", ChangeNickname)
	RegisterRouter("/client/game/profile/nickname/validate", ValidateNickname)
	RegisterRouter("/client/game/profile/nickname/reserved", GetReservedNickname)
	RegisterRouter("/client/profile/status", GetProfileStatus)
}

func CreateProfile(resp http.ResponseWriter, req *http.Request) {

}

func GetProfileData(resp http.ResponseWriter, req *http.Request) {

}

func ReGenerateScav(resp http.ResponseWriter, req *http.Request) {

}

func ChangeVoice(resp http.ResponseWriter, req *http.Request) {

}

func ChangeNickname(resp http.ResponseWriter, req *http.Request) {

}

func ValidateNickname(resp http.ResponseWriter, req *http.Request) {

}

func GetReservedNickname(resp http.ResponseWriter, req *http.Request) {

}

func GetProfileStatus(resp http.ResponseWriter, req *http.Request) {

}
