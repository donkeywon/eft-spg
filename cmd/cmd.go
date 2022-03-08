package cmd

import (
	"github.com/donkeywon/eft-spg/service/cfg"
	"github.com/donkeywon/eft-spg/service/database"
	"github.com/donkeywon/eft-spg/service/eft"
	"github.com/donkeywon/eft-spg/service/httpd"
	"github.com/donkeywon/eft-spg/service/profile"
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

	databaseSvc := database.New(config.Database)
	cfgSvc := cfg.New(config.Cfg)
	httpdSvc := httpd.New(config.Httpd)
	eftSvc := eft.New(config.EFT)
	profileSvc := profile.New(config.Profile)

	c.AppendService(database.Name, databaseSvc)
	c.AppendService(cfg.Name, cfgSvc)
	c.AppendService(httpd.Name, httpdSvc)
	c.AppendService(eft.Name, eftSvc)
	c.AppendService(profile.Name, profileSvc)

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
