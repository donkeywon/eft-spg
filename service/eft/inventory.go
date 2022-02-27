package eft

import (
	"eft-spg/service/database"
	"eft-spg/util"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"go.uber.org/zap"
)

const (
	EquipmentSlotHeadwear            = "Headwear"
	EquipmentSlotEarpiece            = "Earpiece"
	EquipmentSlotFaceCover           = "FaceCover"
	EquipmentSlotArmorVest           = "ArmorVest"
	EquipmentSlotEyewear             = "Eyewear"
	EquipmentSlotArmBand             = "ArmBand"
	EquipmentSlotTacticalVest        = "TacticalVest"
	EquipmentSlotPockets             = "Pockets"
	EquipmentSlotBackpack            = "Backpack"
	EquipmentSlotSecuredContainer    = "SecuredContainer"
	EquipmentSlotFirstPrimaryWeapon  = "FirstPrimaryWeapon"
	EquipmentSlotSecondPrimaryWeapon = "SecondPrimaryWeapon"
	EquipmentSlotHolster             = "Holster"
	EquipmentSlotScabbard            = "Scabbard"
)

func generateInventory(tmplIvt *ast.Node, equipmentChances *ast.Node, generation *ast.Node, botRole string, isPmc bool) {
	ivtBase := generateInventoryBase()

	excludedSlotes := map[string]string{
		EquipmentSlotFirstPrimaryWeapon:  EquipmentSlotFirstPrimaryWeapon,
		EquipmentSlotSecondPrimaryWeapon: EquipmentSlotSecondPrimaryWeapon,
		EquipmentSlotHolster:             EquipmentSlotHolster,
		EquipmentSlotArmorVest:           EquipmentSlotArmorVest,
	}

	tmplIvt.Get("equipment").ForEach(func(path ast.Sequence, node *ast.Node) bool {
		if excludedSlotes[*path.Key] != "" {
			return true
		}

		// TODO generateEquipment
		return true
	})
}

func generateEquipment(ivtBase *ast.Node, equipSlot string, equipPool *ast.Node, modPool *ast.Node, spawnChances *ast.Node, botRole string) *ast.Node {
	spawnChance := int64(100)
	if equipSlot != EquipmentSlotPockets && equipSlot != EquipmentSlotSecuredContainer {
		if !spawnChances.GetByPath("equipment", equipSlot).Exists() {
			return
		}

		spawnChance, _ = spawnChances.GetByPath("equipment", equipSlot).Int64()
	}

	shouldSpawn := util.RandInt(0, 100) <= int(spawnChance)
	if equipPool != nil && equipPool.Exists() && shouldSpawn {
		length, _ := equipPool.Len()
		if length > 0 {
			id := util.GenerateID()
			equipItemTpl, _ := getWeightedInventoryItem(equipPool)
			itemTpl := database.GetDatabase().GetByPath("templates", "items", equipItemTpl)

			if itemTpl == nil || !itemTpl.Exists() {
				svc.Error("Could not find item template", zap.String("tpl", equipItemTpl))
				return nil
			}

			if isItemIncompatibleWithCurrentItem(ivtBase.GetByPath("inventory", "items"), equipItemTpl, equipSlot) {
				return nil
			}

		}
	}

}

func generateExtraPropertiesForItem(itemTpl *ast.Node, botRole string) {
	// TODO
}

func generateInventoryBase() ast.Node {
	equipID := util.GenerateID()
	equipTpl := "55d7217a4bdc2d86028b456d"

	stashID := util.GenerateID()
	stashTpl := "566abbc34bdc2d92178b4576"

	questRaidItemID := util.GenerateID()
	questRaidItemTpl := "5963866286f7747bf429b572"

	questStashItemsID := util.GenerateID()
	questStashItemsTpl := "5963866b86f7747bfa1c4462"

	sortingTableID := util.GenerateID()
	sortingTableTpl := "602543c13fee350cd564d032"

	n, _ := sonic.Get([]byte(fmt.Sprintf(`{
"items": [
    {
        "_id": %s,
        "_tpl": %s
    },
    {
        "_id": %s,
        "_tpl": %s
    },
    {
        "_id": %s,
        "_tpl": %s
    },
    {
        "_id": %s,
        "_tpl": %s
    },
    {
        "_id": %s,
        "_tpl": %s
    }
],
"equipment": %s,
"stash": %s,
"questRaidItems": %s,
"questStashItems": %s,
"sortingTable": %s,
"fastPanel": {}
}`, equipID, equipTpl, stashID, stashTpl, questRaidItemID, questRaidItemTpl, questStashItemsID, questStashItemsTpl,
		sortingTableID, sortingTableTpl, equipID, stashID, questRaidItemID, questStashItemsID, sortingTableID)))

	return n
}
