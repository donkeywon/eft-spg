package database

import (
	"context"
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
	ctx    context.Context
}

func New(config *Config, ctx context.Context) service.Service {
	return &svc{
		config: config,
		ctx:    ctx,
	}
}

func (s *svc) Name() string {
	return Name
}

func (s *svc) Open() error {
	s.logger.Info("Open")

	d, err := util.ReadDatabaseBox()
	if err != nil {
		return err
	}
	s.d = d

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
