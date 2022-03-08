package eft

import (
	"github.com/bytedance/sonic/ast"
	"github.com/donkeywon/eft-spg/service/database"
	"github.com/donkeywon/eft-spg/service/profile"
)

func GetSuits(sessID string) *ast.Node {
	n := ast.NewObject([]ast.Pair{})

	n.Set("_id", *profile.GetProfile(sessID).GetByPath("characters", "pmc", "_id"))
	n.Set("suites", *profile.GetProfile(sessID).GetByPath("suits"))
	return &n
}

func GetTraderSuits(sessID string, traderID string) *ast.Node {
	result := ast.NewArray([]ast.Node{})

	pmcSide, _ := profile.GetPMCProfile(sessID).GetByPath("Info", "Side").String()
	templates := database.GetDatabase().GetByPath("templates", "customization")
	suitesN := database.GetDatabase().GetByPath("traders", traderID, "suits")
	if !suitesN.Exists() {
		return &result
	}
	suites, _ := suitesN.ArrayUseNode()

	for _, suit := range suites {
		suitID, _ := suit.Get("suiteId").String()
		tSuits := templates.Get(suitID)
		if tSuits.Exists() {
			sides, _ := tSuits.GetByPath("_props", "Side").Array()
			for _, side := range sides {
				if side.(string) == pmcSide {
					result.Add(suit)
				}
			}
		}
	}

	return &result
}

func BuyClothing() {
	// TODO
}

func GetAllTraderSuits() {
	// TODO
}
