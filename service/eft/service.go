package eft

import (
	"eft-spg/service/profile"
	"github.com/donkeywon/gtil/service"
)

const (
	Name       = "eft"
	ServerName = "EFT-SPG Server"
	Version    = "0.0.1"
)

func init() {
	profile.ServerVersion = Version
}

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
