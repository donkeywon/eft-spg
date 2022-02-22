package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
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
	s.RegisterRouter("/client/menu/locale/{locale}", s.GetLocalesMenu, true)
	s.RegisterRouter("/client/locale/{locale}", s.GetLocalesGlobal, true)
}

func (s *Svc) GetSettings(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetGlobals(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetTemplateItems(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetTemplateHandbook(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetTemplateSuits(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetTemplateCharacter(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetHideoutProduction(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetHideoutSettings(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetHideoutAreas(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetHideoutScavcase(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetLocalesLanguages(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetLocalesMenu(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}

func (s *Svc) GetLocalesGlobal(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
