package database

import (
	"eft-spg/util"
	"github.com/bytedance/sonic/ast"
	"github.com/donkeywon/gtil/service"
	"go.uber.org/zap"
)

const (
	Name = "database"
)

type svc struct {
	config *Config
	d      *ast.Node
	logger *zap.Logger
}

func New(config *Config) service.Service {
	return &svc{
		config: config,
	}
}

func (s *svc) Name() string {
	return Name
}

func (s *svc) Open() error {
	s.logger.Info("Open")

	d, err := util.ReadJsonDir(s.config.Path)
	if err != nil {
		return err
	}
	s.d = &d

	return err
}

func (s *svc) Close() error {
	s.logger.Info("Close")
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
