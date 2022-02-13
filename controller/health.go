package controller

import "net/http"

func init() {
	RegisterRouter("/player/health/sync", SyncHealth)
}

func SyncHealth(resp http.ResponseWriter, req *http.Request) {

}
