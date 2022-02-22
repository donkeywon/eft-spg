package eft

import (
	"eft-spg/service/cfg"
)

func (s *Svc) GetBotLimit(typ string) (int64, error) {
	if typ == "cursedAssault" || typ == "assaultGroup" {
		typ = "assault"
	}

	return cfg.GetSvc().GetConfig().GetByPath("bot", "presetBatch", typ).Int64()
}
