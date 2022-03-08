package eft

import "github.com/donkeywon/eft-spg/service/database"

const (
	GameEventNone      = "None"
	GameEventChristmas = "Christmas"
	GameEventHalloween = "Halloween"
)

var (
	ChristmasEventItems = []string{
		"5c1a1e3f2e221602b66cc4c2", // White beard
		"5df8a6a186f77412640e2e80", // Red bauble
		"5df8a77486f77412672a1e3f", // Violet bauble
		"5df8a72c86f77412640e2e83", // Silver bauble
		"5a43943586f77416ad2f06e2", // Ded moroz hat
		"5a43957686f7742a2c2f11b0", // Santa hat
	}
)

func ChristmasEventEnabled() bool {
	es, _ := database.GetDatabase().GetByPath("globals", "config", "EventType").Array()
	if len(es) == 0 {
		return false
	}

	for _, e := range es {
		if e.(string) == GameEventChristmas {
			return true
		}
	}

	return false
}

func ItemIsChristmasRelated(itemID string) bool {
	for _, c := range ChristmasEventItems {
		if c == itemID {
			return true
		}
	}

	return false
}
