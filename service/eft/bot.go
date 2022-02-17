package eft

import (
	"github.com/donkeywon/eft-spg/service/cfg"
)

func GetBotLimit(typ string) (int, error) {
	if typ == "cursedAssault" || typ == "assaultGroup" {
		typ = "assault"
	}

	return cfg.Data.MustGet("bot", "presetBatch", typ).Int(), nil
}
