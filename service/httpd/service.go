package httpd

import (
	"compress/zlib"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/donkeywon/eft-spg/util"
	"github.com/donkeywon/gtil/httpd"
	"github.com/donkeywon/gtil/service"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
)

const (
	Name = "httpd"
)

var (
	Methods = []string{"GET", "POST", "PUT"}
	svc     *Svc
)

func GetSvc() *Svc {
	return svc
}

type ServiceProvider func(string, map[string]string, *ast.Node, *http.Request) (interface{}, error)

func (sp ServiceProvider) Handle(sessID string, vars map[string]string, body *ast.Node, r *http.Request) (interface{}, error) {
	return sp(sessID, vars, body, r)
}

type RouteHandler struct {
	sp     ServiceProvider
	logger *zap.Logger
}

func NewRouteHandler(sp ServiceProvider, logger *zap.Logger) *RouteHandler {
	return &RouteHandler{
		sp:     sp,
		logger: logger,
	}
}

func (rh *RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now().UnixNano()
	defer func() {
		end := time.Now().UnixNano()
		rh.logger.Info("Handle req", zap.String("url", r.RequestURI), zap.String("cost", fmt.Sprintf("%.3fms", float64(end-start)/1000000)))
	}()

	err := rh.preHandle(w, r)
	if err != nil {
		rh.logger.Error("Pre handle req fail", zap.Error(err))
	}

	err = rh.handleReq(w, r)
	if err != nil {
		rh.logger.Error("Handle req fail", zap.Error(err))
	}

	err = rh.postHandle(w, r)
	if err != nil {
		rh.logger.Error("Post handle req fail", zap.Error(err))
	}
}

func (rh *RouteHandler) preHandle(w http.ResponseWriter, r *http.Request) error {
	r.RequestURI = strings.Split(r.RequestURI, "?retry=")[0]

	return nil
}

func (rh *RouteHandler) handleReq(w http.ResponseWriter, r *http.Request) error {
	sessID := util.GetSessionID(r)

	var err error
	buf := util.GetBuffer()
	defer buf.Free()
	if r.ContentLength > 0 {
		_, err = buf.ReadFrom(r.Body)
		if err != nil {
			return errors.Wrap(err, util.ErrReadBody)
		}
	}

	var n ast.Node

	if buf.Len() > 0 {
		n, err = sonic.Get(buf.Bytes())
		if err != nil {
			return errors.Wrap(err, util.ErrParseJson)
		}
	} else {
		n = util.GetEmptyJsonNode()
	}

	resp, err := rh.sp.Handle(sessID, mux.Vars(r), &n, r)
	if err != nil {
		rh.logger.Error("Handle req fail", zap.Error(err))
	}

	err = rh.sendResponse(resp, w)
	if err != nil {
		return errors.Wrap(err, util.ErrSendResponse)
	}

	return nil
}

func (rh *RouteHandler) postHandle(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (rh *RouteHandler) sendResponse(resp interface{}, w http.ResponseWriter) error {
	//return util.DoResponseZlibJson(resp, w)
	return util.DoResponseJson(resp, w)
}

type Svc struct {
	*service.BaseService
	httpd  *httpd.HttpD
	Config *httpd.Config

	routers map[string]ServiceProvider

	middleWares []mux.MiddlewareFunc
}

func New(config *httpd.Config) *Svc {
	svc = &Svc{
		BaseService: service.NewBase(),
		httpd:       httpd.New(config),
		Config:      config,
		routers:     make(map[string]ServiceProvider),
		middleWares: []mux.MiddlewareFunc{},
	}

	svc.registerRouter()
	svc.registerMiddleware()

	return svc
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

func (s *Svc) RegisterRouter(path string, sp ServiceProvider) {
	s.routers[path] = sp
}

func (s *Svc) GetRouter() *mux.Router {
	router := mux.NewRouter()

	for path, sp := range s.routers {
		router.Handle(path, NewRouteHandler(sp, s.Logger))
	}

	for _, m := range s.middleWares {
		router.Use(m)
	}

	return router
}

func (s *Svc) sendResponse(resp interface{}, w http.ResponseWriter) error {
	//return util.DoResponseZlibJson(resp, w)
	return util.DoResponseJson(resp, w)
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
	//s.RegisterMiddleware(s.loggingMiddleware)
	s.RegisterMiddleware(s.sessMiddleware)
	//s.RegisterMiddleware(s.decodeMiddleware)
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

func (s *Svc) decodeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			bs := util.GetBuffer()
			defer bs.Free()

			_, err := bs.ReadFrom(r.Body)
			if err != nil {
				s.Error("Read body fail", zap.String("url", r.RequestURI), zap.Error(err))
			}
			err = r.Body.Close()
			if err != nil {
				s.Error("Close body fail", zap.String("url", r.RequestURI), zap.Error(err))
			}

			zr, err := zlib.NewReader(bs)
			r.Body = zr
		}

		// TODO PUT

		next.ServeHTTP(w, r)
	})
}

func (s *Svc) sessMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionID := util.GetSessionID(r)
		if sessionID != "" {
			w.Header().Set("Set-Cookie", "PHPSESSID="+sessionID)
		}

		next.ServeHTTP(w, r)
	})
}

func (s *Svc) responseLogger(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loggingRW := &loggingResponseWriter{
			ResponseWriter: w,
		}
		h.ServeHTTP(loggingRW, r)
		if loggingRW.err != nil {
			s.Error("Response fail", zap.String("url", r.RequestURI), zap.Int("code", loggingRW.statusCode), zap.Error(loggingRW.err))
		}
	}
}

type loggingResponseWriter struct {
	http.ResponseWriter

	statusCode int
	err        error
}

func (w *loggingResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *loggingResponseWriter) Write(body []byte) (int, error) {
	i, err := w.ResponseWriter.Write(body)
	if err != nil {
		w.err = err
	}
	return i, err
}
