package cfg

import (
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"github.com/donkeywon/eft-spg/service/cfg/hook"
	"github.com/donkeywon/eft-spg/util"
	"github.com/donkeywon/gtil/service"
)

const (
	Name = "cfg"
)

var (
	Data = jsonvalue.NewObject()
)

type svc struct {
	*service.BaseService
	config *Config
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
	c, err := util.ReadConfigBox()
	if err != nil {
		return err
	}
	err = hook.PostLoadHook(c)
	if err != nil {
		return err
	}

	Data = c

	return nil
}

func (s *svc) Close() error {
	return nil
}

func (s *svc) Shutdown() error {
	return nil
}
