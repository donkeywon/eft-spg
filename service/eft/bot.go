package eft

import (
	"eft-spg/service/cfg"
	"eft-spg/util"
	"github.com/pkg/errors"
)

func (s *Svc) GetBotLimit(typ string) (int64, error) {
	if typ == "" {
		return 0, errors.Errorf(util.ErrIllegalArg, typ)
	}

	if typ == "cursedAssault" || typ == "assaultGroup" {
		typ = "assault"
	}

	return cfg.GetSvc().GetConfig().GetByPath("bot", "presetBatch", typ).Int64()
}
