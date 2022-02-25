package httpd

import (
	"eft-spg/service/database"
	"eft-spg/util"
	"github.com/bytedance/sonic/ast"
	"net/http"
	"strings"
	"time"
)

func (s *Svc) registerDataRouter() {
	s.RegisterRouter("/client/settings", s.GetSettings, true)
	s.RegisterRouter("/client/globals", s.GetGlobals, true)
	s.RegisterRouter("/client/items", s.GetTemplateItems, true)
	s.RegisterRouter("/client/handbook/templates", s.GetTemplateHandbook, true)
	s.RegisterRouter("/client/customization", s.GetTemplateSuits, true)
	s.RegisterRouter("/client/account/customization", s.GetTemplateCharacter, true)
	s.RegisterRouter("/client/hideout/production/recipes", s.GetHideoutProduction, true)
	s.RegisterRouter("/client/hideout/settings", s.GetHideoutSettings, true)
	s.RegisterRouter("/client/hideout/areas", s.GetHideoutAreas, true)
	s.RegisterRouter("/client/hideout/production/scavcase/recipes", s.GetHideoutScavcase, true)
	s.RegisterRouter("/client/languages", s.GetLocalesLanguages, true)
	s.RegisterRouter("/client/menu/locale/", s.GetLocalesMenu, false)
	s.RegisterRouter("/client/locale/", s.GetLocalesGlobal, false)
}

func (s *Svc) GetSettings(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return database.GetSvc().GetDatabase().Get("settings"), nil
}

func (s *Svc) GetGlobals(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	g := database.GetSvc().GetDatabase().Get("globals")
	g.SetAny("time", time.Now().Unix())
	return g, nil
}

func (s *Svc) GetTemplateItems(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	i := database.GetSvc().GetDatabase().GetByPath("templates", "items")
	return util.GetResponseWrapperFromData(i), nil
}

func (s *Svc) GetTemplateHandbook(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	b := database.GetSvc().GetDatabase().GetByPath("templates", "handbook")
	return util.GetResponseWrapperFromData(b), nil
}

func (s *Svc) GetTemplateSuits(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	c := database.GetSvc().GetDatabase().GetByPath("templates", "customization")
	return util.GetResponseWrapperFromData(c), nil
}

func (s *Svc) GetTemplateCharacter(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	c := database.GetSvc().GetDatabase().GetByPath("templates", "character")
	return util.GetResponseWrapperFromData(c), nil
}

func (s *Svc) GetTemplateQuests(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	q := database.GetSvc().GetDatabase().GetByPath("templates", "quests")
	return util.GetResponseWrapperFromData(q), nil
}

func (s *Svc) GetHideoutProduction(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	p := database.GetSvc().GetDatabase().GetByPath("hideout", "production")
	return util.GetResponseWrapperFromData(p), nil
}

func (s *Svc) GetHideoutSettings(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	st := database.GetSvc().GetDatabase().GetByPath("hideout", "settings")
	return util.GetResponseWrapperFromData(st), nil
}

func (s *Svc) GetHideoutAreas(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	a := database.GetSvc().GetDatabase().GetByPath("hideout", "areas")
	return util.GetResponseWrapperFromData(a), nil
}

func (s *Svc) GetHideoutScavcase(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	sc := database.GetSvc().GetDatabase().GetByPath("hideout", "scavcase")
	return util.GetResponseWrapperFromData(sc), nil
}

func (s *Svc) GetLocalesLanguages(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	l := database.GetSvc().GetDatabase().GetByPath("locales", "languages")
	return util.GetResponseWrapperFromData(l), nil
}

func (s *Svc) GetLocalesMenu(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	m := database.GetSvc().GetDatabase().GetByPath("locales", "menu", strings.Split(r.RequestURI, "/client/menu/locale/")[1])
	return util.GetResponseWrapperFromData(m), nil
}

func (s *Svc) GetLocalesGlobal(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	g := database.GetSvc().GetDatabase().GetByPath("locales", "global", strings.Split(r.RequestURI, "/client/locale/")[1])
	return util.GetResponseWrapperFromData(g), nil
}
