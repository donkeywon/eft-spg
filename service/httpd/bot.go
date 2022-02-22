package httpd

import (
	"eft-spg/service/eft"
	"github.com/bytedance/sonic/ast"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

func (s *Svc) registerBotRouter() {
	s.RegisterRouter("/singleplayer/settings/bot/limit/", s.GetBotLimit, false)
	s.RegisterRouter("/singleplayer/settings/bot/difficulty/", s.GetBotDifficulty, false)
	s.RegisterRouter("/client/game/bot/generate", s.GenerateBots, true)
	s.RegisterRouter("/singleplayer/settings/bot/maxCap", s.GetBotCap, false)
}

func (s *Svc) GetBotLimit(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	ss := strings.Split(r.RequestURI, "/")

	b, err := eft.GetSvc().GetBotLimit(ss[len(ss)-1])
	if err != nil {
		return nil, errors.Wrap(err, "Get bot limit fail")
	}

	return b, nil
}

func (s *Svc) GetBotDifficulty(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GenerateBots(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetBotCap(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
