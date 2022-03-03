package httpd

import (
	"eft-spg/service/database"
	"eft-spg/util"
	"github.com/bytedance/sonic/ast"
	"net/http"
	"strconv"
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
	return util.GetResponseWrapperFromData(database.GetDatabase().Get("settings")), nil
}

func (s *Svc) GetGlobals(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	g := database.GetDatabase().Get("globals")
	g.Set("time", ast.NewString(strconv.Itoa(int(time.Now().Unix()))))
	return g, nil
}

func (s *Svc) GetTemplateItems(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	i := database.GetDatabase().GetByPath("templates", "items")
	return util.GetResponseWrapperFromData(i), nil
}

func (s *Svc) GetTemplateHandbook(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	b := database.GetDatabase().GetByPath("templates", "handbook")
	return util.GetResponseWrapperFromData(b), nil
}

func (s *Svc) GetTemplateSuits(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	c := database.GetDatabase().GetByPath("templates", "customization")
	return util.GetResponseWrapperFromData(c), nil
}

func (s *Svc) GetTemplateCharacter(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	c := database.GetDatabase().GetByPath("templates", "character")
	return util.GetResponseWrapperFromData(c), nil
}

func (s *Svc) GetTemplateQuests(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	q := database.GetDatabase().GetByPath("templates", "quests")
	return util.GetResponseWrapperFromData(q), nil
}

func (s *Svc) GetHideoutProduction(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	p := database.GetDatabase().GetByPath("hideout", "production")
	return util.GetResponseWrapperFromData(p), nil
}

func (s *Svc) GetHideoutSettings(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	st := database.GetDatabase().GetByPath("hideout", "settings")
	return util.GetResponseWrapperFromData(st), nil
}

func (s *Svc) GetHideoutAreas(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	a := database.GetDatabase().GetByPath("hideout", "areas")
	return util.GetResponseWrapperFromData(a), nil
}

func (s *Svc) GetHideoutScavcase(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	sc := database.GetDatabase().GetByPath("hideout", "scavcase")
	return util.GetResponseWrapperFromData(sc), nil
}

func (s *Svc) GetLocalesLanguages(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	l := database.GetDatabase().GetByPath("locales", "languages")
	return util.GetResponseWrapperFromData(l), nil
}

func (s *Svc) GetLocalesMenu(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	m := database.GetDatabase().GetByPath("locales", "menu", vars["locale"])
	return util.GetResponseWrapperFromData(m), nil
}

func (s *Svc) GetLocalesGlobal(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	g := database.GetDatabase().GetByPath("locales", "global", vars["locale"])
	return util.GetResponseWrapperFromData(g), nil
}
