package cmd

import (
	"github.com/donkeywon/eft-spg/service/cfg"
	"github.com/donkeywon/eft-spg/service/httpd"
	"github.com/donkeywon/gtil/service"
)

const (
	Name = "cmd"
)

type Command struct {
	*service.BaseService
	config *Config
}

func New(config *Config) *Command {
	c := &Command{
		BaseService: service.NewBase(),
		config:      config,
	}

	c.AppendService(httpd.Name, httpd.New(config.Httpd))
	//c.AppendService(database.Name, database.New(config.Database))
	c.AppendService(cfg.Name, cfg.New(config.Cfg))

	return c
}

func (c *Command) Name() string {
	return Name
}

func (c *Command) Open() error {
	return nil
}

func (c *Command) Close() error {
	return nil
}

func (c *Command) Shutdown() error {
	return nil
}
