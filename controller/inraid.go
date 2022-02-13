package controller

import "net/http"

func init() {
	RegisterRouter("/raid/profile/save", SaveProfile)
	RegisterRouter("/singleplayer/settings/raid/endstate", GetRaidEndState)
	RegisterRouter("/singleplayer/settings/weapon/durability", GetWeaponDurability)
	RegisterRouter("/singleplayer/settings/raid/menu", GetRaidMenuSettings)
}

func SaveProfile(resp http.ResponseWriter, req *http.Request) {

}

func GetRaidEndState(resp http.ResponseWriter, req *http.Request) {

}

func GetWeaponDurability(resp http.ResponseWriter, req *http.Request) {

}

func GetRaidMenuSettings(resp http.ResponseWriter, req *http.Request) {

}

func RegisterPlayer(resp http.ResponseWriter, req *http.Request) {

}
