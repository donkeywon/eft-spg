package eft

import (
	"eft-spg/service/cfg"
)

func GetBotLimit(typ string) (int64, error) {
	if typ == "cursedAssault" || typ == "assaultGroup" {
		typ = "assault"
	}

	return cfg.GetConfig().GetByPath("bot", "presetBatch", typ).Int64()
}
