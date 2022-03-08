package eft

import (
	"github.com/bytedance/sonic/ast"
	"github.com/donkeywon/eft-spg/service/database"
	"github.com/donkeywon/eft-spg/util"
)

func isItemIncompatibleWithCurrentItem(items *ast.Node, tplIDToCheck string, equSlot string) bool {
	// TODO: Can probably be optimized to cache itemTemplates as items are added to inventory
	itemTpls := ast.NewArray(nil)
	items.ForEach(func(path ast.Sequence, item *ast.Node) bool {
		tplID, _ := item.Get("_tpl").String()
		tpl := database.GetDatabase().GetByPath("templates", "items", tplID)
		if tpl.Exists() {
			itemTpls.Add(*item)
		}
		return true
	})

	tplToCheck := database.GetDatabase().GetByPath("templates", "items", tplIDToCheck)

	// Check if any of the current inventory templates have the incoming item defined as incompatible
	curIvtCheck := false
	itemTpls.ForEach(func(path ast.Sequence, itemTpl *ast.Node) bool {
		blockEquSlot := false
		if itemTpl.GetByPath("_props", "Blocks"+equSlot).Exists() {
			blockEquSlot, _ = itemTpl.GetByPath("_props", "Blocks"+equSlot).Bool()
		}
		conflictItems, _ := itemTpl.GetByPath("_props", "ConflictingItems").Array()

		if blockEquSlot || util.ArrayContains(conflictItems, equSlot) {
			curIvtCheck = true
			return false
		}

		return true
	})

	// Check if the incoming item has any inventory items defined as incompatible
	itemCheck := false
	items.ForEach(func(path ast.Sequence, item *ast.Node) bool {
		itemSlotID, _ := item.Get("slotId").String()
		blockEquSlot := false
		if tplToCheck.GetByPath("_props", "Blocks"+itemSlotID).Exists() {
			blockEquSlot, _ = tplToCheck.GetByPath("_props", "Blocks"+itemSlotID).Bool()
		}
		conflictItems, _ := tplToCheck.GetByPath("_props", "ConflictingItems").Array()

		if blockEquSlot || util.ArrayContains(conflictItems, equSlot) {
			itemCheck = true
			return false
		}

		return true
	})

	return curIvtCheck || itemCheck
}
