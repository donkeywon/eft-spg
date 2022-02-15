package hook

import (
	"github.com/bytedance/sonic/ast"
	"github.com/donkeywon/eft-spg/helper"
)

func init() {
	RegisterHook(BotHook)
}

var whiteList = []string{
	helper.ItemBaseClass["Jewelry"],
	helper.ItemBaseClass["Electronics"],
	helper.ItemBaseClass["BuildingMaterial"],
	helper.ItemBaseClass["Tool"],
	helper.ItemBaseClass["HouseholdGoods"],
	helper.ItemBaseClass["MedicalSupplies"],
	helper.ItemBaseClass["Lubricant"],
	helper.ItemBaseClass["Battery"],
	helper.ItemBaseClass["Keycard"],
	helper.ItemBaseClass["KeyMechanical"],
	helper.ItemBaseClass["AssaultScope"],
	helper.ItemBaseClass["ReflexSight"],
	helper.ItemBaseClass["TacticalCombo"],
	helper.ItemBaseClass["Magazine"],
	helper.ItemBaseClass["Knife"],
	helper.ItemBaseClass["BarterItem"],
	helper.ItemBaseClass["Silencer"],
	helper.ItemBaseClass["Foregrip"],
	helper.ItemBaseClass["Info"],
	helper.ItemBaseClass["Food"],
	helper.ItemBaseClass["Drink"],
	helper.ItemBaseClass["Drugs"],
	helper.ItemBaseClass["Armor"],
	helper.ItemBaseClass["Stimulator"],
	helper.ItemBaseClass["AmmoBox"],
	helper.ItemBaseClass["Money"],
	helper.ItemBaseClass["Other"],
}

func BotHook(n *ast.Node) error {
	_, err := n.GetByPath("bot", "pmc", "dynamicLoot").SetAny("whitelist", whiteList)
	return err
}
