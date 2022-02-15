package database

import (
	"github.com/donkeywon/eft-spg/util"
	"github.com/donkeywon/gtil/service"
)

const (
	Name = "database"
)

type svc struct {
	*service.BaseService
	config *Config
	d      []byte
}

func New(config *Config) service.Service {
	return &svc{
		BaseService: service.NewBase(),
		config:      config,
	}
}

func (s *svc) Name() string {
	return Name
}

func (s *svc) Open() error {
	d, err := util.ReadDatabaseBox()
	if err != nil {
		return err
	}
	s.d = d

	return err
}

func (s *svc) Close() error {
	return nil
}

func (s *svc) Shutdown() error {
	return nil
}
