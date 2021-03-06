package cmd

import (
	"github.com/donkeywon/eft-spg/service/cfg"
	"github.com/donkeywon/eft-spg/service/database"
	"github.com/donkeywon/eft-spg/service/eft"
	"github.com/donkeywon/eft-spg/service/profile"
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
	EFT      *eft.Config      `json:"eft" yaml:"eft"`
}

func NewConfig() *Config {
	return &Config{
		Cfg:      cfg.NewConfig(),
		Database: database.NewConfig(),
		Httpd:    httpd.NewConfig(DefaultAddr),
		Profile:  profile.NewConfig(),
		EFT:      eft.NewConfig(),
	}
}
