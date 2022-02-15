package httpd

import (
	"context"
	"eft-spg/controller"
	"github.com/donkeywon/gtil/httpd"
	"github.com/donkeywon/gtil/service"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

const (
	Name = "httpd"
)

type svc struct {
	logger *zap.Logger
	httpd  *httpd.HttpD
	ctx    context.Context
}

func New(config *httpd.Config, ctx context.Context) service.Service {
	s := &svc{
		ctx: ctx,
	}
	s.httpd = httpd.New(config, context.Background())
	s.httpd.SetHandler(controller.GetRouter())

	return s
}

func (s *svc) Name() string {
	return Name
}

func (s *svc) Open() error {
	s.logger.Info("Open")
	return multierr.Combine(s.httpd.Open(), s.httpd.LastError())
}

func (s *svc) Close() error {
	s.logger.Info("Close")
	return s.httpd.Close()
}

func (s *svc) Shutdown() error {
	s.logger.Info("Shutdown")
	return s.httpd.Shutdown()
}

func (s *svc) WithLogger(logger *zap.Logger) {
	s.logger = logger.Named(Name)
}

func (s *svc) Statistics() map[string]float64 {
	return nil
}

func (s *svc) LastError() error {
	return nil
}
