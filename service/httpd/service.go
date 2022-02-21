package httpd

import (
	"eft-spg/util"
	"fmt"
	"github.com/donkeywon/gtil/httpd"
	"github.com/donkeywon/gtil/service"
	"github.com/gorilla/mux"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"net/http"
)

const (
	Name       = "httpd"
	ServerName = "EFT-SPG Server"
)

var (
	Methods = []string{"GET", "POST", "PUT"}
)

type Svc struct {
	*service.BaseService
	httpd  *httpd.HttpD
	Config *httpd.Config

	routers map[string]func(http.ResponseWriter, *http.Request)

	middleWares []mux.MiddlewareFunc
}

func New(config *httpd.Config) *Svc {
	s := &Svc{
		BaseService: service.NewBase(),
		httpd:       httpd.New(config),
		Config:      config,
		routers:     make(map[string]func(http.ResponseWriter, *http.Request)),
		middleWares: []mux.MiddlewareFunc{},
	}

	s.registerRouter()

	return s
}

func (s *Svc) Name() string {
	return Name
}

func (s *Svc) Open() error {
	s.Info("Opening")
	s.httpd.SetHandler(s.GetRouter())
	err := multierr.Combine(s.httpd.Open(), s.httpd.LastError())
	if err != nil {
		return err
	}
	s.Info("Opened", zap.String("addr", s.Config.Addr))
	return nil
}

func (s *Svc) Close() error {
	return s.httpd.Close()
}

func (s *Svc) Shutdown() error {
	return s.httpd.Shutdown()
}

func (s *Svc) RegisterMiddleware(middlewareFunc mux.MiddlewareFunc) {
	s.middleWares = append(s.middleWares, middlewareFunc)
}

func (s *Svc) RegisterRouter(path string, f func(http.ResponseWriter, *http.Request)) {
	s.routers[path] = f
}

func (s *Svc) GetRouter() *mux.Router {
	router := mux.NewRouter()

	for path, f := range s.routers {
		router.HandleFunc(path, f).Methods(Methods...)
	}

	for _, m := range s.middleWares {
		router.Use(m)
	}

	return router
}

func (s *Svc) registerRouter() {
	s.registerBotRouter()
	s.registerBundleRouter()
	s.registerCustomizationRouter()
	s.registerDataRouter()
	s.registerDialogRouter()
	s.registerGameRouter()
	s.registerHealthRouter()
	s.registerImageRouter()
	s.registerInraidRouter()
	s.registerInsuranceRouter()
	s.registerItemEventRouter()
	s.registerLauncherRouter()
	s.registerLocationRouter()
	s.registerMatchRouter()
	s.registerNotifierRouter()
	s.registerPresetBuildRouter()
	s.registerProfileRouter()
	s.registerQuestRouter()
	s.registerRagfairRouter()
	s.registerTraderRouter()
	s.registerWeatherRouter()
}

func (s *Svc) registerMiddleware() {
	s.RegisterMiddleware(s.loggingMiddleware)
	s.RegisterMiddleware(s.authMiddleware)
}

func (s *Svc) backendUrl() string {
	return fmt.Sprintf("http://%s", s.Config.Addr)
}

func (s *Svc) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Info("Handle req", zap.String("url", r.RequestURI))

		next.ServeHTTP(w, r)
	})
}

func (s *Svc) authMiddleware(next http.Handler) http.Handler {
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

func (s *Svc) logRespErr(err error, r *http.Request) {
	if err != nil {
		s.Error("Response fail", zap.String("url", r.RequestURI), zap.Error(err))
	}
}
