package cmd

import (
	"eft-spg/service/cfg"
	"eft-spg/service/database"
	"eft-spg/service/profile"
	"github.com/donkeywon/gtil/httpd"
)

const (
	DefaultAddr = "127.0.0.1:1408"
)

type Config struct {
	Cfg      *cfg.Config      `json:"cfg" yaml:"cfg"`
	Database *database.Config `json:"database" yaml:"database"`
	Httpd    *httpd.Config    `json:"httpd" yaml:"httpd"`
	Profile  *profile.Config  `json:"profile" yaml:"profile"`
}

func NewConfig() *Config {
	return &Config{
		Cfg:      cfg.NewConfig(),
		Database: database.NewConfig(),
		Httpd:    httpd.NewConfig(DefaultAddr),
	}
}
