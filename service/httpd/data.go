package httpd

import (
	"net/http"
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

func (s *Svc) GetSettings(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetGlobals(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetTemplateItems(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetTemplateHandbook(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetTemplateSuits(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetTemplateCharacter(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetHideoutProduction(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetHideoutSettings(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetHideoutAreas(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetHideoutScavcase(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetLocalesLanguages(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetLocalesMenu(resp http.ResponseWriter, req *http.Request) {

}

func (s *Svc) GetLocalesGlobal(resp http.ResponseWriter, req *http.Request) {

}
