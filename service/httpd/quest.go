package httpd

import (
	"net/http"
)

func (s *Svc) registerQuestRouter() {
	s.RegisterRouter("/client/quest/list", s.ListQuests)
	s.RegisterRouter("/client/repeatalbeQuests/activityPeriods", s.ActivityPeriods)
}

func (s *Svc) ListQuests(resp http.ResponseWriter, req *http.Request) {}

func (s *Svc) ActivityPeriods(resp http.ResponseWriter, req *http.Request) {}
