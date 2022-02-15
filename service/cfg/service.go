package cfg

import (
	"context"
	"eft-spg/service/cfg/hook"
	"eft-spg/util"
	"github.com/bytedance/sonic/ast"
	"github.com/donkeywon/gtil/service"
	"go.uber.org/zap"
)

const (
	Name = "cfg"
)

type svc struct {
	config *Config
	logger *zap.Logger
	c      *ast.Node
	ctx    context.Context
	cancel context.CancelFunc
}

func New(config *Config, ctx context.Context) service.Service {
	return &svc{
		config: config,
	}
}

func (s *svc) Name() string {
	return Name
}

func (s *svc) Open() error {
	s.logger.Info("Open")

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
	s.logger.Info("Close")
	s.cancel()
	return nil
}

func (s *svc) Shutdown() error {
	s.logger.Info("Shutdown")
	return nil
}

func (s *svc) WithLogger(logger *zap.Logger) {
	s.logger = logger.Named(s.Name())
}

func (s *svc) Statistics() map[string]float64 {
	return nil
}

func (s *svc) LastError() error {
	return nil
}

func (s *svc) Get(key string) *ast.Node {
	return s.c.Get(key)
}

func (s *svc) GetByPath(path ...interface{}) *ast.Node {
	return s.c.GetByPath(path...)
}
