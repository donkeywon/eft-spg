package eft

import (
	"eft-spg/service/database"
	"eft-spg/util"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"go.uber.org/zap"
    "strings"
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

func generateEquipment(ivtBase *ast.Node, equipSlot string, equipPool *ast.Node, modPool *ast.Node, spawnChances *ast.Node, botRole string) ast.Node {
    items := ast.NewArray(nil)

	spawnChance := int64(100)
	if equipSlot != EquipmentSlotPockets && equipSlot != EquipmentSlotSecuredContainer {
		if !spawnChances.GetByPath("equipment", equipSlot).Exists() {
			return items
		}

		spawnChance, _ = spawnChances.GetByPath("equipment", equipSlot).Int64()
	}

	shouldSpawn := util.RandInt(0, 100) <= int(spawnChance)
    if equipPool == nil || equipPool.Exists() || !shouldSpawn {
        return items
    }

    length, _ := equipPool.Len()
    if length <= 0 {
        return items
    }

    equipItemTpl, _ := getWeightedInventoryItem(equipPool)
    itemTpl := database.GetDatabase().GetByPath("templates", "items", equipItemTpl)

    if itemTpl == nil || !itemTpl.Exists() {
        svc.Error("Could not find item template", zap.String("tpl", equipItemTpl))
        return items
    }

    if isItemIncompatibleWithCurrentItem(ivtBase.GetByPath("inventory", "items"), equipItemTpl, equipSlot) {
        return items
    }

    id := util.GenerateID()
    item := ast.NewObject([]ast.Pair{{
        Key: "_id",
        Value: ast.NewString(id),
    }, {
        Key: "_tpl",
        Value: ast.NewString(equipItemTpl),
    }, {
        Key: "parentId",
        Value: *ivtBase.Get("equipment"),
    }, {
        Key: "slotId",
        Value: ast.NewString(equipSlot),
    }})

    extraPropsForItem := generateExtraPropertiesForItem(itemTpl, botRole)
    extraPropsForItem.ForEach(func(path ast.Sequence, node *ast.Node) bool {
        item.Set(*path.Key, *node)
    })

    hasMod := false
    modPool.ForEach(func(path ast.Sequence, node *ast.Node) bool {
        if *path.Key == equipItemTpl {
            hasMod = true
            items = generateModsForItem(ast.NewArray([]ast.Node{item}), modPool, id, itemTpl, spawnChances.Get("mods"), false)
        }
        return true
    })
    if !hasMod {
        items.Add(item)
    }

    return items
}

var ammoContainers = map[string]bool{
    "mod_magazine": true,
    "patron_in_weapon": true,
    "patron_in_weapon_000": true,
    "patron_in_weapon_001": true,
    "cartridges": true,
}
func generateModsForItem(items ast.Node, modPool *ast.Node, parentID string, parentTpl *ast.Node, modSpawnChances *ast.Node, isPmc bool) ast.Node {
    mis := ast.NewArray(nil)
    parentTplId, _ := parentTpl.Get("_id").String()
    itemModPool := modPool.Get(parentTplId)

    slotsLength, _ := parentTpl.GetByPath("_props", "Slots").Len()
    cartridgesLength, _ := parentTpl.GetByPath("_props", "Cartridges").Len()
    chambersLength, _ := parentTpl.GetByPath("_props", "Chambers").Len()
    if slotsLength == 0 && cartridgesLength == 0 && chambersLength == 0 {
        svc.Error("No slots to support item", zap.String("parentTplID", parentTplId))
        return mis
    }

    itemModPool.ForEach(func(path ast.Sequence, node *ast.Node) bool {
        var itemSlot *ast.Node

        modSlot, _ := node.String()
        switch modSlot {
        case "patron_in_weapon":
        case "patron_in_weapon_000":
        case "patron_in_weapon_001":
            parentTpl.GetByPath("_props", "Chambers").ForEach(func(path ast.Sequence, node *ast.Node) bool {
                n, _ := node.Get("_name").String()
                if strings.Index(n, modSlot) != -1 {
                    itemSlot = node
                }
                return true
            })
            break
        case "cartridges":
            parentTpl.GetByPath("_props", "Cartridges").ForEach(func(path ast.Sequence, node *ast.Node) bool {
                n, _ := node.Get("_name").String()
                if n == modSlot {
                    itemSlot = node
                }
            })
            break
        default:
            parentTpl.GetByPath("_props", "Slots").ForEach(func(path ast.Sequence, node *ast.Node) bool {
                n, _ := node.Get("_name").String()
                if n == modSlot {
                    itemSlot = node
                }
            })
            break
        }

        if itemSlot == nil {
            svc.Error("Slot does not exist for item", zap.String("slot", modSlot), zap.String("parentTplID", parentTplId))
            return true
        }

        modSpawnChance, _ := modSpawnChances.Get(modSlot).Int64()
        itemSlotRequired, _ := itemSlot.Get("_required").Bool()
        if itemSlotRequired || ammoContainers[modSlot] == true {
            modSpawnChance = 100
        }
        if util.RandInt(0, 100) > int(modSpawnChance) {
            return true
        }

        // TODO
    })
}

func generateExtraPropertiesForItem(itemTpl *ast.Node, botRole string) ast.Node {
	props := ast.NewObject(nil)

	if itemTpl.GetByPath("_props", "MaxDurability").Exists() {
        var maxDurability, currentDurability int
		if itemTpl.GetByPath("_props", "weapClass").Exists() {
			maxDurability = getRandomisedMaxWeaponOrArmorDurability(botRole, "weapon")
			currentDurability = getRandomisedWeaponOrArmorDurability(botRole, "weapon", maxDurability)
		} else if itemTpl.GetByPath("_props", "armorClass").Exists() {
			armorClass, _ := itemTpl.GetByPath("_props", "armorClass").Int64()
			if int(armorClass) == 0 {
				m, _ := itemTpl.GetByPath("_props", "MaxDurability").Int64()
				c, _ := itemTpl.GetByPath("_props", "MaxDurability").Int64()
                maxDurability = int(m)
                currentDurability = int(c)
			} else {
				maxDurability = getRandomisedMaxWeaponOrArmorDurability(botRole, "armor"))
				currentDurability = getRandomisedWeaponOrArmorDurability(botRole, "armor", maxDurability))
			}
		}

        repairable := ast.NewObject([]ast.Pair{{
            Key:   "Durability",
            Value: ast.NewNumber(string(currentDurability)),
        }, {
            Key:   "MaxDurability",
            Value: ast.NewNumber(string(maxDurability)),
        }})

        props.Set("Repairable", repairable)
	}

    if itemTpl.GetByPath("_props", "HasHinge").Exists() {
        props.Set("Togglable", ast.NewObject([]ast.Pair{{Key: "On", Value: ast.NewBool(true)}}))
    }
    if itemTpl.GetByPath("_props", "Foldable").Exists() {
        props.Set("Foldable", ast.NewObject([]ast.Pair{{Key: "Folded", Value: ast.NewBool(false)}}))
    }
    if itemTpl.GetByPath("_props", "weapFireType").Exists() {
        ft, _ := itemTpl.GetByPath("_props", "weapFireType").Array()
        if len(ft) > 0 {
            props.SetAny("FireMode", ast.NewObject([]ast.Pair{{Key: "FireMode", Value: ast.NewString(util.RandChoose(ft).(string))}}))
        }
    }
    if itemTpl.GetByPath("_props", "MaxHpResource").Exists() {
        props.Set("Medkit", ast.NewObject([]ast.Pair{{Key: "HpResource", Value: *itemTpl.GetByPath("_props", "MaxHpResource")}}))
    }
    if itemTpl.GetByPath("_props", "MaxResource").Exists() && itemTpl.GetByPath("_props", "foodUseTime").Exists() {
        props.Set("FoodDrink", ast.NewObject([]ast.Pair{{Key: "HpPercent", Value: *itemTpl.GetByPath("_props", "MaxResource")}}))
    }

    l, _ := props.Len()
    if l > 0 {
        return ast.NewObject([]ast.Pair{{Key: "upd", Value: props}})
    } else {
        return props
    }
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
