package httpd

import (
	"eft-spg/service/database"
	"eft-spg/util"
	"github.com/bytedance/sonic/ast"
	"net/http"
	"time"
)

func (s *Svc) registerDataRouter() {
	s.RegisterRouter("/client/settings", s.GetSettings)
	s.RegisterRouter("/client/globals", s.GetGlobals)
	s.RegisterRouter("/client/items", s.GetTemplateItems)
	s.RegisterRouter("/client/handbook/templates", s.GetTemplateHandbook)
	s.RegisterRouter("/client/customization", s.GetTemplateSuits)
	s.RegisterRouter("/client/account/customization", s.GetTemplateCharacter)
	s.RegisterRouter("/client/hideout/production/recipes", s.GetHideoutProduction)
	s.RegisterRouter("/client/hideout/settings", s.GetHideoutSettings)
	s.RegisterRouter("/client/hideout/areas", s.GetHideoutAreas)
	s.RegisterRouter("/client/hideout/production/scavcase/recipes", s.GetHideoutScavcase)
	s.RegisterRouter("/client/languages", s.GetLocalesLanguages)
	s.RegisterRouter("/client/menu/locale/{locale}", s.GetLocalesMenu)
	s.RegisterRouter("/client/locale/{locale}", s.GetLocalesGlobal)
}

func (s *Svc) GetSettings(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return util.GetResponseWrapperFromData(database.GetSvc().GetDatabase().Get("settings")), nil
}

func (s *Svc) GetGlobals(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	g := database.GetSvc().GetDatabase().Get("globals")
	g.SetAny("time", time.Now().Unix())
	return g, nil
}

func (s *Svc) GetTemplateItems(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	i := database.GetSvc().GetDatabase().GetByPath("templates", "items")
	return util.GetResponseWrapperFromData(i), nil
}

func (s *Svc) GetTemplateHandbook(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	b := database.GetSvc().GetDatabase().GetByPath("templates", "handbook")
	return util.GetResponseWrapperFromData(b), nil
}

func (s *Svc) GetTemplateSuits(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	c := database.GetSvc().GetDatabase().GetByPath("templates", "customization")
	return util.GetResponseWrapperFromData(c), nil
}

func (s *Svc) GetTemplateCharacter(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	c := database.GetSvc().GetDatabase().GetByPath("templates", "character")
	return util.GetResponseWrapperFromData(c), nil
}

func (s *Svc) GetTemplateQuests(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	q := database.GetSvc().GetDatabase().GetByPath("templates", "quests")
	return util.GetResponseWrapperFromData(q), nil
}

func (s *Svc) GetHideoutProduction(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	p := database.GetSvc().GetDatabase().GetByPath("hideout", "production")
	return util.GetResponseWrapperFromData(p), nil
}

func (s *Svc) GetHideoutSettings(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	st := database.GetSvc().GetDatabase().GetByPath("hideout", "settings")
	return util.GetResponseWrapperFromData(st), nil
}

func (s *Svc) GetHideoutAreas(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	a := database.GetSvc().GetDatabase().GetByPath("hideout", "areas")
	return util.GetResponseWrapperFromData(a), nil
}

func (s *Svc) GetHideoutScavcase(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	sc := database.GetSvc().GetDatabase().GetByPath("hideout", "scavcase")
	return util.GetResponseWrapperFromData(sc), nil
}

func (s *Svc) GetLocalesLanguages(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	l := database.GetSvc().GetDatabase().GetByPath("locales", "languages")
	return util.GetResponseWrapperFromData(l), nil
}

func (s *Svc) GetLocalesMenu(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	m := database.GetSvc().GetDatabase().GetByPath("locales", "menu", vars["locale"])
	return util.GetResponseWrapperFromData(m), nil
}

func (s *Svc) GetLocalesGlobal(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	g := database.GetSvc().GetDatabase().GetByPath("locales", "global", vars["locale"])
	return util.GetResponseWrapperFromData(g), nil
}
