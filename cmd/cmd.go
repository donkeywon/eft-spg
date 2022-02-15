package cmd

import (
	"context"
	"eft-spg/service/cfg"
	"eft-spg/service/database"
	"eft-spg/service/httpd"
	"github.com/donkeywon/gtil/service"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

const (
	Name = "cmd"
)

type Command struct {
	logger *zap.Logger
	ctx    context.Context
	cancel context.CancelFunc
	config *Config

	httpd    service.Service
	database service.Service
	cfg      service.Service
}

func New(config *Config, ctx context.Context) *Command {
	ctx1, cancel := context.WithCancel(ctx)
	return &Command{
		config:   config,
		ctx:      ctx1,
		cancel:   cancel,
		httpd:    httpd.New(config.Httpd, ctx1),
		database: database.New(config.Database, ctx1),
		cfg:      cfg.New(config.Cfg, ctx1),
	}
}

func (c *Command) Name() string {
	return Name
}

func (c *Command) Open() error {
	c.logger.Info("Open")
	return multierr.Combine(c.database.Open(), c.cfg.Open(), c.httpd.Open())
}

func (c *Command) Close() error {
	c.logger.Info("Close")
	c.cancel()
}

func (c *Command) Shutdown() error {
	c.logger.Info("Shutdown")
	//TODO implement me
	panic("implement me")
}

func (c *Command) WithLogger(logger *zap.Logger) {
	c.logger = logger.Named(Name)
}

func (c *Command) Statistics() map[string]float64 {
	return nil
}
