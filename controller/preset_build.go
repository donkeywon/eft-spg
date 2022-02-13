package controller

import "net/http"

func init() {
	RegisterRouter("/client/handbook/builds/my/list", GetHandbookUserList)
}

func GetHandbookUserList(resp http.ResponseWriter, req *http.Request) {

}
