package controller

import (
	"github.com/donkeywon/eft-spg/util"
	//_ "github.com/donkeywon/eft-spg/controller/middleware"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func init() {
	RegisterMiddleware(loggingMiddleware)
	RegisterMiddleware(authMiddleware)
}

const (
	Name = "controller"
)

var (
	routers = make(map[string]func(http.ResponseWriter, *http.Request))

	middleWares []mux.MiddlewareFunc

	methods = []string{"GET", "POST", "PUT"}

	Logger *zap.Logger
)

func RegisterMiddleware(middlewareFunc mux.MiddlewareFunc) {
	middleWares = append(middleWares, middlewareFunc)
}

func RegisterRouter(path string, f func(http.ResponseWriter, *http.Request)) {
	routers[path] = f
}

func GetRouter() *mux.Router {
	router := mux.NewRouter()

	for path, f := range routers {
		router.HandleFunc(path, f).Methods(methods...)
	}

	for _, m := range middleWares {
		router.Use(m)
	}

	return router
}

func WithLogger(l *zap.Logger) {
	Logger = l.Named(Name)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Logger.Info("Handle req", zap.String("url", r.RequestURI))

		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionID := util.GetSessionID(r)
		if sessionID == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		w.Header().Set("Set-Cookie", "PHPSESSID="+sessionID)
		next.ServeHTTP(w, r)
	})
}
