package httpd

import (
	"eft-spg/controller"
	"github.com/donkeywon/gtil/httpd"
	"github.com/donkeywon/gtil/service"
	"go.uber.org/multierr"
)

const (
	Name = "httpd"
)

type svc struct {
	*service.BaseService
	httpd *httpd.HttpD
}

func New(config *httpd.Config) service.Service {
	s := &svc{
		BaseService: service.NewBase(),
		httpd:       httpd.New(config),
	}
	s.httpd.SetHandler(controller.GetRouter())

	return s
}

func (s *svc) Name() string {
	return Name
}

func (s *svc) Open() error {
	return multierr.Combine(s.httpd.Open(), s.httpd.LastError())
}

func (s *svc) Close() error {
	return s.httpd.Close()
}

func (s *svc) Shutdown() error {
	return s.httpd.Shutdown()
}
