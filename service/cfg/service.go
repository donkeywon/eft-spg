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

type svc struct {
	*service.BaseService
	config *Config
	c      *ast.Node
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

	s.c = c

	return nil
}

func (s *svc) Close() error {
	return nil
}

func (s *svc) Shutdown() error {
	return nil
}

func (s *svc) Get(key string) *ast.Node {
	return s.c.Get(key)
}

func (s *svc) GetByPath(path ...interface{}) *ast.Node {
	return s.c.GetByPath(path...)
}
