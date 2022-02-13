package controller

import "net/http"

func init() {
	RegisterRouter("/client/settings", GetSettings)
	RegisterRouter("/client/globals", GetGlobals)
	RegisterRouter("/client/items", GetTemplateItems)
	RegisterRouter("/client/handbook/templates", GetTemplateHandbook)
	RegisterRouter("/client/customization", GetTemplateSuits)
	RegisterRouter("/client/account/customization", GetTemplateCharacter)
	RegisterRouter("/client/hideout/production/recipes", GetHideoutProduction)
	RegisterRouter("/client/hideout/settings", GetHideoutSettings)
	RegisterRouter("/client/hideout/areas", GetHideoutAreas)
	RegisterRouter("/client/hideout/production/scavcase/recipes", GetHideoutScavcase)
	RegisterRouter("/client/languages", GetLocalesLanguages)
	RegisterRouter("/client/menu/locale/{locale}", GetLocalesMenu)
	RegisterRouter("/client/locale/{locale}", GetLocalesGlobal)
}

func GetSettings(resp http.ResponseWriter, req *http.Request) {

}

func GetGlobals(resp http.ResponseWriter, req *http.Request) {

}

func GetTemplateItems(resp http.ResponseWriter, req *http.Request) {

}

func GetTemplateHandbook(resp http.ResponseWriter, req *http.Request) {

}

func GetTemplateSuits(resp http.ResponseWriter, req *http.Request) {

}

func GetTemplateCharacter(resp http.ResponseWriter, req *http.Request) {

}

func GetHideoutProduction(resp http.ResponseWriter, req *http.Request) {

}

func GetHideoutSettings(resp http.ResponseWriter, req *http.Request) {

}

func GetHideoutAreas(resp http.ResponseWriter, req *http.Request) {

}

func GetHideoutScavcase(resp http.ResponseWriter, req *http.Request) {

}

func GetLocalesLanguages(resp http.ResponseWriter, req *http.Request) {

}

func GetLocalesMenu(resp http.ResponseWriter, req *http.Request) {

}

func GetLocalesGlobal(resp http.ResponseWriter, req *http.Request) {

}
