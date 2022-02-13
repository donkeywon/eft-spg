package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

var (
	routers = make(map[string]func(http.ResponseWriter, *http.Request))

	methods = []string{"GET", "POST", "PUT"}
)

func RegisterRouter(path string, f func(http.ResponseWriter, *http.Request)) {
	routers[path] = f
}

func GetRouter() *mux.Router {
	router := mux.NewRouter()

	for path, f := range routers {
		router.HandleFunc(path, f).Methods(methods...)
	}

	return router
}