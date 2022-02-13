package controller

import "net/http"

func init() {
	RegisterRouter("/singleplayer/bundles", GetBundles)
	RegisterRouter(".bundle", GetBundles)
}

func GetBundles(resp http.ResponseWriter, req *http.Request) {

}
