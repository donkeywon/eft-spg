package cmd

import (
	"eft-spg/service/database"
	"github.com/donkeywon/gtil/httpd"
)

type Config struct {
	Database *database.Config `json:"database" yaml:"database"`
	Httpd    *httpd.Config    `json:"httpd" yaml:"httpd"`
}
