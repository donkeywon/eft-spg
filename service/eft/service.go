package eft

import (
	"github.com/donkeywon/gtil/service"
)

const (
	Name = "eft"
)

var (
	svc *Svc
)

func GetSvc() *Svc {
	return svc
}

type Svc struct {
	*service.BaseService
	config *Config
}

func New(config *Config) *Svc {
	svc = &Svc{
		BaseService: service.NewBase(),
		config:      config,
	}
	return svc
}

func (s *Svc) Name() string {
	return Name
}

func (s *Svc) Open() error {
	return nil
}

func (s *Svc) Close() error {
	return nil
}

func (s *Svc) Shutdown() error {
	return nil
}
