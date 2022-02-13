package httpd

import (
	"context"
	"eft-spg/controller"
	"github.com/donkeywon/gtil/httpd"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

const (
	Name = "httpd"
)

type Service struct {
	logger *zap.Logger
	httpd  *httpd.HttpD
}

func New(config *httpd.Config) *Service {
	s := &Service{}
	s.httpd = httpd.New(config, context.Background())
	s.httpd.SetHandler(controller.GetRouter())

	return s
}

func (s *Service) Name() string {
	return Name
}

func (s *Service) Open() error {
	s.logger.Info("Open")
	return multierr.Combine(s.httpd.Open(), s.httpd.LastError())
}

func (s *Service) Close() error {
	s.logger.Info("Close")
	return s.httpd.Close()
}

func (s *Service) Shutdown() error {
	s.logger.Info("Shutdown")
	return s.httpd.Shutdown()
}

func (s *Service) WithLogger(logger *zap.Logger) {
	s.logger = logger.Named(Name)
}

func (s *Service) Statistics() map[string]float64 {
	return nil
}

func (s *Service) LastError() error {
	return nil
}
