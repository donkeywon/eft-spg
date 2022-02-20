package httpd

import (
	"net/http"
)

func (s *Svc) registerItemEventRouter() {
	s.RegisterRouter("/client/game/profile/items/moving", s.HandleEvents)
}

func (s *Svc) HandleEvents(resp http.ResponseWriter, req *http.Request) {

}
