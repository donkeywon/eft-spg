package cmd

import (
	"eft-spg/service/database"
	"github.com/donkeywon/gtil/httpd"
)

type Config struct {
	Cfg      *config.Config   `json:"cfg" yaml:"cfg"`
	Database *database.Config `json:"database" yaml:"database"`
	Httpd    *httpd.Config    `json:"httpd" yaml:"httpd"`
}
