package cfg

import (
	"eft-spg/service/cfg/hook"
	"eft-spg/util"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"github.com/donkeywon/gtil/service"
)

const (
	Name = "cfg"
)

var (
	cfg = jsonvalue.NewObject()
)

type Svc struct {
	*service.BaseService
	Config *Config
}

func New(config *Config) *Svc {
	return &Svc{
		BaseService: service.NewBase(),
		Config:      config,
	}
}

func (s *Svc) Name() string {
	return Name
}

func (s *Svc) Open() error {
	c, err := util.ReadConfigBox()
	if err != nil {
		return err
	}
	err = hook.PostLoadHook(c)
	if err != nil {
		return err
	}

	cfg = c

	return nil
}

func (s *Svc) Close() error {
	return nil
}

func (s *Svc) Shutdown() error {
	return nil
}

func GetConfig() *jsonvalue.V {
	return cfg
}
