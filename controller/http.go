package controller

import "net/http"

func init() {
	RegisterRouter(".jpg", GetImage)
	RegisterRouter(".png", GetImage)
	RegisterRouter(".ico", GetImage)
}

func GetImage(resp http.ResponseWriter, req *http.Request) {

}
