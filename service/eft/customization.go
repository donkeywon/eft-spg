package eft

import (
	"eft-spg/service/database"
	"eft-spg/service/profile"
	"github.com/bytedance/sonic/ast"
)

func (s *Svc) GetSuits(sessID string) *ast.Node {
	n := ast.NewObject([]ast.Pair{})

	n.Set("_id", *profile.GetSvc().GetProfile(sessID).GetByPath("characters", "pmc", "_id"))
	n.Set("suites", *profile.GetSvc().GetProfile(sessID).GetByPath("suits"))
	return &n
}

func (s *Svc) GetTraderSuits(sessID string, traderID string) *ast.Node {
	result := ast.NewArray([]ast.Node{})

	pmcSide, _ := profile.GetSvc().GetPMCProfile(sessID).GetByPath("Info", "Side").String()
	templates := database.GetSvc().GetDatabase().GetByPath("templates", "customization")
	suitesN := database.GetSvc().GetDatabase().GetByPath("traders", traderID, "suits")
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

func (s *Svc) BuyClothing() {
	// TODO
}

func (s *Svc) GetAllTraderSuits() {
	// TODO
}
