package httpd

import (
	"eft-spg/service/eft"
	"github.com/bytedance/sonic/ast"
	"github.com/pkg/errors"
	"net/http"
)

func (s *Svc) registerBotRouter() {
	s.RegisterRouter("/singleplayer/settings/bot/limit/{type}", s.GetBotLimit)
	s.RegisterRouter("/singleplayer/settings/bot/difficulty/{type}/{difficulty}", s.GetBotDifficulty)
	s.RegisterRouter("/client/game/bot/generate", s.GenerateBots)
	s.RegisterRouter("/singleplayer/settings/bot/maxCap", s.GetBotCap)
}

func (s *Svc) GetBotLimit(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	b, err := eft.GetBotLimit(vars["type"])
	if err != nil {
		return nil, errors.Wrap(err, "Get bot limit fail")
	}

	return b, nil
}

func (s *Svc) GetBotDifficulty(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return eft.GetBotDifficulty(vars["type"], vars["difficulty"])
}

func (s *Svc) GenerateBots(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetBotCap(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return eft.GetBotCap(), nil
}
