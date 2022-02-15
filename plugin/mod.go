package main

import (
	"github.com/buger/jsonparser"
	"github.com/donkeywon/eft-spg/service/cfg"
)

func GetValue() {
	cfg.Data, _ = jsonparser.Set(cfg.Data, []byte(`119`), "bot", "presetBatch", "assault")
}
