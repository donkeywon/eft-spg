package eft

import (
	"eft-spg/service/cfg"
)

func GetBotLimit(typ string) (int, error) {
	if typ == "cursedAssault" || typ == "assaultGroup" {
		typ = "assault"
	}

	return cfg.GetConfig().MustGet("bot", "presetBatch", typ).Int(), nil
}
