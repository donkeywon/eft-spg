package controller

import (
	"github.com/donkeywon/eft-spg/service/eft"
	"github.com/donkeywon/eft-spg/util"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func init() {
	RegisterRouter("/singleplayer/settings/bot/limit/{type}", GetBotLimit)
	RegisterRouter("/singleplayer/settings/bot/difficulty/{type}/{difficulty}", GetBotDifficulty)
	RegisterRouter("/client/game/bot/generate", GenerateBots)
	RegisterRouter("/singleplayer/settings/bot/maxCap", GetBotCap)
}

func GetBotLimit(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	b, err := eft.GetBotLimit(vars["type"])
	if err != nil {
		Logger.Error("GetBotLimit fail", zap.Error(err))
	}

	err = util.DoResponseJson(b, resp)
	if err != nil {
		Logger.Error("Response fail", zap.Error(err))
	}
}

func GetBotDifficulty(resp http.ResponseWriter, req *http.Request) {

}

func GenerateBots(resp http.ResponseWriter, req *http.Request) {

}

func GetBotCap(resp http.ResponseWriter, req *http.Request) {

}
