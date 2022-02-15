package service

import (
	"github.com/buger/jsonparser"
	"github.com/donkeywon/eft-spg/service/cfg"
)

func GetBotLimit(typ string) (int64, error) {
	if typ == "cursedAssault" || typ == "assaultGroup" {
		typ = "assault"
	}

	return jsonparser.GetInt(cfg.Data, "bot", "presetBatch", typ)
}
