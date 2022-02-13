package controller

import "net/http"

func init() {
	RegisterRouter("/client/quest/list", ListQuests)
	RegisterRouter("/client/repeatalbeQuests/activityPeriods", ActivityPeriods)
}

func ListQuests(resp http.ResponseWriter, req *http.Request) {}

func ActivityPeriods(resp http.ResponseWriter, req *http.Request) {}
