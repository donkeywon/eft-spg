package controller

import (
	"net/http"
)

func init() {
	RegisterRouter("/singleplayer/settings/bot/limit/{type}", GetBotLimit)
	RegisterRouter("/singleplayer/settings/bot/difficulty/{type}/{difficulty}", GetBotDifficulty)
	RegisterRouter("/client/game/bot/generate", GenerateBots)
	RegisterRouter("/singleplayer/settings/bot/maxCap", GetBotCap)
}

func GetBotLimit(resp http.ResponseWriter, req *http.Request) {

}

func GetBotDifficulty(resp http.ResponseWriter, req *http.Request) {

}

func GenerateBots(resp http.ResponseWriter, req *http.Request) {

}

func GetBotCap(resp http.ResponseWriter, req *http.Request) {

}
