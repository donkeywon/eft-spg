package controller

import "net/http"

func init() {
	RegisterRouter("/client/game/profile/items/moving", HandleEvents)
}

func HandleEvents(resp http.ResponseWriter, req *http.Request) {

}
