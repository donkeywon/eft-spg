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
	config *Config
	cfg    *ast.Node
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

func GetCfg() *ast.Node {
	return svc.cfg
}
