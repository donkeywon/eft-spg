package eft

import (
	"eft-spg/service/cfg"
	"github.com/donkeywon/gtil/service"
)

const (
	Name = "eft"
)

var (
	svc *Svc
)

func GetSvc() *Svc {
	return svc
}

type Svc struct {
	*service.BaseService
	config *Config
}

func New(config *Config) *Svc {
	svc = &Svc{
		BaseService: service.NewBase(),
		config:      config,
	}
	return svc
}

func (s *Svc) Name() string {
	return Name
}

func (s *Svc) Open() error {
	s.initVar()
	return nil
}

func (s *Svc) Close() error {
	return nil
}

func (s *Svc) Shutdown() error {
	return nil
}

func (s *Svc) initVar() {
	BotRoleBear, _ = cfg.GetCfg().GetByPath("bot", "pmc", "bearType").String()
	BotRoleUsec, _ = cfg.GetCfg().GetByPath("bot", "pmc", "bearType").String()

	bosses, _ := cfg.GetCfg().GetByPath("bot", "bosses").Array()
	BotRoleBoss = make([]string, len(bosses))
	for i, boss := range bosses {
		BotRoleBoss[i] = boss.(string)
	}
}
