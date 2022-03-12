package eft

import (
	"github.com/bytedance/sonic/ast"
	"github.com/donkeywon/eft-spg/service/database"
	"github.com/donkeywon/eft-spg/util"
	"go.uber.org/zap"
)

func generateWeapon(ivtBase *ast.Node, equipSlot string, weaponPool *ast.Node, modPool *ast.Node, modChances *ast.Node, magCounts *ast.Node, botRole string, isPmc bool) {
	id := util.GenerateID()
	weaponTplId, _ := getWeightedInventoryItem(weaponPool)
	itemTpl := database.GetDatabase().GetByPath("templates", "items", weaponTplId)

	if !itemTpl.Exists() {
		svc.Error("Cound not find item template", zap.String("tplId", weaponTplId), zap.String("weaponSlot", equipSlot))
		return
	}

	tmpMod := ast.NewObject(nil)
	tmpMod.Set("_id", ast.NewString(id))
	tmpMod.Set("_tpl", ast.NewString(weaponTplId))
	tmpMod.Set("parentId", *ivtBase.GetByPath("inventory", "equipment"))
	tmpMod.Set("slotId", ast.NewString(equipSlot))
	extraPropsForItem := generateExtraPropertiesForItem(itemTpl, botRole)
	extraPropsForItem.ForEach(func(path ast.Sequence, node *ast.Node) bool {
		tmpMod.Set(*path.Key, *node)
		return true
	})
	weaponMods := ast.NewArray([]ast.Node{tmpMod})
	if util.ArrayContainsString(util.JsonNodeKeys(modPool), weaponTplId) {
		generateModsForItem(&weaponMods, modPool, id, itemTpl, modChances, isPmc)
	}
}

func isWeaponValid(itemList *ast.Node) bool {
	err1 := itemList.ForEach(func(path ast.Sequence, item *ast.Node) bool {
		itemTplID, _ := item.Get("_tpl").String()
		itemID, _ := item.Get("_id").String()
		tpl := database.GetDatabase().GetByPath("templates", "items", itemTplID)
		tplSlots := tpl.GetByPath("_props", "Slots")
		if !tplSlots.Exists() {
			return true
		}

		tplSlots.Load()
		tplSlotsLen, _ := tplSlots.Len()
		if tplSlotsLen == 0 {
			return true
		}

		err2 := tplSlots.ForEach(func(path ast.Sequence, slotNode *ast.Node) bool {
			sr, _ := slotNode.Get("_required").Bool()
			if !sr {
				return true
			}

			found := false
			slotName, _ := slotNode.Get("_name").String()
			itemList.ForEach(func(path ast.Sequence, it *ast.Node) bool {
				parentId, _ := it.Get("_parentId").String()
				slotId, _ := it.Get("_slotId").String()
				if parentId == itemID && slotId == slotName {
					found = true
					return false
				}
				return true
			})

			if !found {
				tplId, _ := tpl.Get("_id").String()
				svc.Error("Required slot was empty", zap.String("slotName", slotName), zap.String("tplId", tplId))
				return false
			}
			return found
		})
	})
}
