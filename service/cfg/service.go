package cfg

import (
	"eft-spg/service/cfg/hook"
	"eft-spg/util"
	"github.com/bytedance/sonic/ast"
	"github.com/donkeywon/gtil/service"
)

const (
	Name = "cfg"
)

var (
	svc *Svc
)

func GetSvc() *Svc {
	return svc
}

type Svc struct {
	*service.BaseService
	Config *Config
	cfg    *ast.Node
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
	svc = s
	c, err := util.ReadConfigBox()
	if err != nil {
		return err
	}
	err = hook.PostLoadHook(c)
	if err != nil {
		return err
	}

	s.cfg = c

	return nil
}

func (s *Svc) Close() error {
	return nil
}

func (s *Svc) Shutdown() error {
	return nil
}

func (s *Svc) GetConfig() *ast.Node {
	return s.cfg
}
