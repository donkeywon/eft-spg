package httpd

import (
	"eft-spg/service/eft"
	"eft-spg/util"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func (s *Svc) registerBotRouter() {
	s.RegisterRouter("/singleplayer/settings/bot/limit/{type}", s.GetBotLimit)
	s.RegisterRouter("/singleplayer/settings/bot/difficulty/{type}/{difficulty}", s.GetBotDifficulty)
	s.RegisterRouter("/client/game/bot/generate", s.GenerateBots)
	s.RegisterRouter("/singleplayer/settings/bot/maxCap", s.GetBotCap)
}

func (s *Svc) GetBotLimit(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	b, err := eft.GetSvc().GetBotLimit(vars["type"])
	if err != nil {
		s.Error("GetBotLimit fail", zap.Error(err))
		return
	}

	err = util.DoResponseJson(b, resp)
	if err != nil {
		s.Error("Response fail", zap.Error(err))
		return
	}
}

func (s *Svc) GetBotDifficulty(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GenerateBots(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetBotCap(resp http.ResponseWriter, req *http.Request) {

}
