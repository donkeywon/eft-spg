package eft

import (
	"github.com/bytedance/sonic/ast"
	"github.com/donkeywon/eft-spg/service/cfg"
	"github.com/donkeywon/eft-spg/service/database"
	"github.com/donkeywon/eft-spg/util"
	"go.uber.org/zap"
	"sort"
	"strings"
)

const (
	EquipSlotHeadwear            = "Headwear"
	EquipSlotEarpiece            = "Earpiece"
	EquipSlotFaceCover           = "FaceCover"
	EquipSlotArmorVest           = "ArmorVest"
	EquipSlotEyewear             = "Eyewear"
	EquipSlotArmBand             = "ArmBand"
	EquipSlotTacticalVest        = "TacticalVest"
	EquipSlotPockets             = "Pockets"
	EquipSlotBackpack            = "Backpack"
	EquipSlotSecuredContainer    = "SecuredContainer"
	EquipSlotFirstPrimaryWeapon  = "FirstPrimaryWeapon"
	EquipSlotSecondPrimaryWeapon = "SecondPrimaryWeapon"
	EquipSlotHolster             = "Holster"
	EquipSlotScabbard            = "Scabbard"
)

var excludedSlotes = map[string]string{
	EquipSlotFirstPrimaryWeapon:  EquipSlotFirstPrimaryWeapon,
	EquipSlotSecondPrimaryWeapon: EquipSlotSecondPrimaryWeapon,
	EquipSlotHolster:             EquipSlotHolster,
	EquipSlotArmorVest:           EquipSlotArmorVest,
}

func generateInventory(tmplIvt *ast.Node, equipChances *ast.Node, generation *ast.Node, botRole string, isPmc bool) {
	// Generate base inventory with no items
	ivtBase := generateInventoryBase()

	// Go over all defined equipment slots and generate an item for each of them

	tmplIvt.Get("equipment").ForEach(func(path ast.Sequence, equip *ast.Node) bool {
		// Weapons have special generation and will be generated seperately; ArmorVest should be generated after TactivalVest
		equipSlot := *path.Key
		if excludedSlotes[equipSlot] == "" {
			generateEquipment(ivtBase, equipSlot, equip, tmplIvt.Get("mods"), equipChances, botRole)
		}
		return true
	})

	// ArmorVest is generated afterwards to ensure that TacticalVest is always first, in case it is incompatible
	generateEquipment(ivtBase, EquipSlotArmorVest, tmplIvt.GetByPath("equipment", "ArmorVest"), tmplIvt.Get("mods"), equipChances, botRole)

	// Roll weapon spawns and generate a weapon for each roll that passed
	firstPrimaryWeaponChance, _ := equipChances.GetByPath("equipment", EquipSlotFirstPrimaryWeapon).Int64()
	secondPrimaryWeaponChance, _ := equipChances.GetByPath("equipment", EquipSlotSecondPrimaryWeapon).Int64()
	holsterChance, _ := equipChances.GetByPath("equipment", EquipSlotHolster).Int64()
	shouldSpawnPrimary := util.RandInt(0, 100) <= int(firstPrimaryWeaponChance)
	shouldSpawnSecond := false
	if shouldSpawnPrimary {
		// only roll for a chance at secondary if primary roll was successful
		shouldSpawnSecond = util.RandInt(0, 100) <= int(secondPrimaryWeaponChance)
	}
	shouldSpawnHolster := true
	if shouldSpawnPrimary {
		// roll for an extra pistol, unless primary roll failed - in that case, pistol is guaranteed
		shouldSpawnHolster = util.RandInt(0, 100) <= int(holsterChance)
	}

	primarySpawn := ast.NewObject(nil)
	primarySpawn.Set("slot", ast.NewString(EquipSlotFirstPrimaryWeapon))
	primarySpawn.Set("shouldSpawn", ast.NewBool(shouldSpawnPrimary))
	secondSpawn := ast.NewObject(nil)
	secondSpawn.Set("slot", ast.NewString(EquipSlotSecondPrimaryWeapon))
	secondSpawn.Set("shouldSpawn", ast.NewBool(shouldSpawnSecond))
	holsterSpawn := ast.NewObject(nil)
	holsterSpawn.Set("slot", ast.NewString(EquipSlotHolster))
	holsterSpawn.Set("shouldSpawn", ast.NewBool(shouldSpawnHolster))
	//weaponSlotSpawns := ast.NewArray([]ast.Node{primarySpawn, secondSpawn, holsterSpawn})
	// TODO generateWeapon
}

func generateWeapon() {

}

func generateEquipment(ivtBase *ast.Node, equipSlot string, equipPool *ast.Node, modPool *ast.Node, spawnChances *ast.Node, botRole string) *ast.Node {
	it := ast.NewArray(nil)
	items := &it

	spawnChance := int64(100)
	if equipSlot != EquipSlotPockets && equipSlot != EquipSlotSecuredContainer {
		if !spawnChances.GetByPath("equipment", equipSlot).Exists() {
			return items
		}

		spawnChance, _ = spawnChances.GetByPath("equipment", equipSlot).Int64()
	}

	shouldSpawn := util.RandInt(0, 100) <= int(spawnChance)
	if equipPool == nil || equipPool.Exists() || !shouldSpawn {
		return items
	}

	equipPool.Load()
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
		// Bad luck - randomly picked item was not compatible with current gear
		return items
	}

	id := util.GenerateID()
	item := ast.NewObject(nil)
	item.Set("_id", ast.NewString(id))
	item.Set("_tpl", ast.NewString(equipItemTpl))
	item.Set("parentId", *ivtBase.Get("equipment"))
	item.Set("slotId", ast.NewString(equipSlot))

	extraPropsForItem := generateExtraPropertiesForItem(itemTpl, botRole)
	extraPropsForItem.ForEach(func(path ast.Sequence, node *ast.Node) bool {
		item.Set(*path.Key, *node)
		return true
	})

	items.Add(item)
	modPool.ForEach(func(path ast.Sequence, node *ast.Node) bool {
		if *path.Key == equipItemTpl {
			items = generateModsForItem(items, modPool, id, itemTpl, spawnChances.Get("mods"), false)
			return false
		}
		return true
	})

	items.ForEach(func(path ast.Sequence, i *ast.Node) bool {
		ivtBase.Get("items").Add(*i)
		return true
	})

	return items
}

var ammoContainers = map[string]bool{
	"mod_magazine":         true,
	"patron_in_weapon":     true,
	"patron_in_weapon_000": true,
	"patron_in_weapon_001": true,
	"cartridges":           true,
}

func generateModsForItem(items *ast.Node, modPool *ast.Node, parentID string, parentTpl *ast.Node, modSpawnChances *ast.Node, isPmc bool) *ast.Node {
	mis := ast.NewArray(nil)
	parentTplId, _ := parentTpl.Get("_id").String()
	itemModPool := modPool.Get(parentTplId)

	slotsLength, _ := parentTpl.GetByPath("_props", "Slots").Len()
	cartridgesLength, _ := parentTpl.GetByPath("_props", "Cartridges").Len()
	chambersLength, _ := parentTpl.GetByPath("_props", "Chambers").Len()
	if slotsLength == 0 && cartridgesLength == 0 && chambersLength == 0 {
		svc.Error("No slots to support item", zap.String("parentTplID", parentTplId))
		return &mis
	}

	cartridgeBlacklist, _ := cfg.GetCfg().GetByPath("bot", "pmc", "cartridgeBlacklist").Array()

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

				return true
			})
			break
		default:
			parentTpl.GetByPath("_props", "Slots").ForEach(func(path ast.Sequence, node *ast.Node) bool {
				n, _ := node.Get("_name").String()
				if n == modSlot {
					itemSlot = node
				}

				return true
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

		// Filter blacklisted cartridges
		if isPmc && ammoContainers[modSlot] == true {
			// Array includes mod_magazine which isnt a cartridge, but we need to filter the other 4 items
			itemModPoolModSlot := itemModPool.Get(modSlot)
			itemModPoolModSlot.ForEach(func(path ast.Sequence, node *ast.Node) bool {
				id, _ := node.String()
				if util.ArrayContains(cartridgeBlacklist, id) {
					itemModPoolModSlot.UnsetByIndex(path.Index)
				}
				return true
			})
		}

		found := false
		modTplID := ""
		itemModTpls, _ := itemModPool.Get(modSlot).Array()
		for len(itemModTpls) > 0 {
			randIdx := util.RandInt(0, len(itemModTpls))
			modTplID = itemModTpls[randIdx].(string)
			if !isItemIncompatibleWithCurrentItem(items, modTplID, modSlot) {
				found = true
				break
			}

			itemModTpls = append(itemModTpls[:randIdx], itemModTpls[randIdx+1:]...)
		}

		// Find a mod to attach from items db for required slots if none found above
		var parentSlot *ast.Node
		parentTpl.GetByPath("_props", "Slots").ForEach(func(path ast.Sequence, slot *ast.Node) bool {
			slotName, _ := slot.Get("_name").String()
			if slotName == modSlot {
				parentSlot = slot
				return false
			}
			return true
		})
		if !found && parentSlot != nil {
			parentSlotRequired, _ := parentSlot.Get("_required").Bool()
			if parentSlotRequired {
				modTplID = getModTplIDFromItemDB(modTplID, parentSlot, modSlot, items)
				if modTplID != "" {
					found = true
				}
			}
		}

		if !found || modTplID == "" {
			if itemSlotRequired {
				svc.Error("Could not locate any compatible items to fill",
					zap.String("modSlot", modSlot), zap.String("parentTplID", parentTplId))
			}
			return true
		}

		itemSlotFilter, _ := itemSlot.GetByPath("_props", "filters", 0, "Filter").Array()
		if !util.ArrayContains(itemSlotFilter, modTplID) {
			svc.Error("Mod is not compatible with slot for item",
				zap.String("modTplID", modTplID), zap.String("modSlot", modSlot), zap.String("parentTplID", parentTplId))
			return true
		}

		modTpl := database.GetDatabase().GetByPath("templates", "items", modTplID)
		if !modTpl.Exists() {
			svc.Error("Could not find mod item template with tpl",
				zap.String("modTplID", modTplID), zap.String("parentTplID", parentTplId), zap.String("modSlot", modSlot))
			return true
		}

		// TODO: check if weapon already has sight
		// 'sight' 550aa4154bdc2dd8348b456b 2x parents down
		parentItemID, _ := modTpl.Get("_parent").String()
		parentItem := database.GetDatabase().GetByPath("templates", "items", parentItemID)
		parentParentItemID, _ := parentItem.Get("_parent").String()
		if parentItemID == "550aa4154bdc2dd8348b456b" || parentParentItemID == "550aa4154bdc2dd8348b456b" {
			// TODO, check if another sight is already on gun AND isnt a side-mounted sight
			// if weapon has sight already, skip
		}

		modID := util.GenerateID()
		item := ast.NewObject(nil)
		item.Set("_id", ast.NewString(modID))
		item.Set("_tpl", ast.NewString(modTplID))
		item.Set("parentId", ast.NewString(parentID))
		item.Set("slotId", ast.NewString(modSlot))
		extraProps := generateExtraPropertiesForItem(modTpl, "")
		extraProps.ForEach(func(path ast.Sequence, node *ast.Node) bool {
			item.Set(*path.Key, *node)
			return true
		})
		items.Add(item)

		// I first thought we could use the recursive generateModsForItems as previously for cylinder magazines.
		// However, the recurse doesnt go over the slots of the parent mod but over the modPool which is given by the bot config
		// where we decided to keep cartridges instead of camoras. And since a CylinderMagazine only has one cartridge entry and
		// this entry is not to be filled, we need a special handling for the CylinderMagazine
		parentItemName, _ := parentItem.Get("_name").String()
		if parentItemName == "CylinderMagazine" {
			// we don't have child mods, we need to create the camoras for the magazines instead
			fillCamora(items, modPool, modID, modTpl)
		} else {
			modPool.ForEach(func(path ast.Sequence, node *ast.Node) bool {
				if *path.Key == modTplID {
					generateModsForItem(items, modPool, modID, modTpl, modSpawnChances, false)
					return false
				}

				return true
			})
		}
		return true
	})

	return items
}

/**
 * With the shotgun revolver (60db29ce99594040e04c4a27) 12.12 introduced CylinderMagazines.
 * Those magazines (e.g. 60dc519adf4c47305f6d410d) have a "Cartridges" entry with a _max_count=0.
 * Ammo is not put into the magazine directly but assigned to the magazine's slots: The "camora_xxx" slots.
 * This function is a helper called by generateModsForItem for mods with parent type "CylinderMagazine"
 *
 * @param items               The items where the CylinderMagazine's camora are appended to
 * @param modPool             modPool which should include available cartrigdes
 * @param parentId            The CylinderMagazine's UID
 * @param parentTemplate      The CylinderMagazine's template
 */
func fillCamora(items *ast.Node, modPool *ast.Node, parentID string, parentTpl *ast.Node) {
	parentTplID, _ := parentTpl.GetByPath("_id").String()
	itemModPool := modPool.Get(parentTplID)

	modSlot := "cartridges"
	if !itemModPool.Get(modSlot).Exists() {
		svc.Error("itemPool does not contain cartridges for a CylinderMagazine. Filling of camoras cancelled.", zap.String("parentTplID", parentTplID))
		return
	}

	modTplID := ""
	found := false
	itemModTpls, _ := itemModPool.Get(modSlot).Array()
	for len(itemModTpls) > 0 {
		randIdx := util.RandInt(0, len(itemModTpls))
		modTplID = itemModTpls[randIdx].(string)
		if !isItemIncompatibleWithCurrentItem(items, modTplID, modSlot) {
			found = true
			break
		}

		itemModTpls = append(itemModTpls[:randIdx], itemModTpls[randIdx+1:]...)
	}

	if !found {
		svc.Error("No compatible ammo found. Filling of camoras cancelled.", zap.String("modSlot", modSlot))
		return
	}

	parentTpl.GetByPath("_props", "Slots").ForEach(func(path ast.Sequence, slot *ast.Node) bool {
		modSlot, _ = slot.Get("_name").String()
		modID := util.GenerateID()
		item := ast.NewObject(nil)
		item.Set("_id", ast.NewString(modID))
		item.Set("_tpl", ast.NewString(modTplID))
		item.Set("parentId", ast.NewString(parentID))
		item.Set("slotId", ast.NewString(modSlot))
		items.Add(item)
		return true
	})
}

type sortModArr []string

func (ma sortModArr) Len() int {
	return len(ma)
}

func (ma sortModArr) Swap(i, j int) {
	ma[i], ma[j] = ma[j], ma[i]
}

func (ma sortModArr) Less(i, j int) bool {
	isc, _ := database.GetDatabase().GetByPath("templates", "items", ma[i], "_props", "SpawnChance").Int64()
	jsc, _ := database.GetDatabase().GetByPath("templates", "items", ma[j], "_props", "SpawnChance").Int64()

	return isc < jsc
}

func getModTplIDFromItemDB(modTpl string, parentSlot *ast.Node, modSlot string, items *ast.Node) string {
	// Find compatible mods and make an array of them
	var modArr []string
	parentSlot.GetByPath("_props", "filters", 0, "Filter").ForEach(func(path ast.Sequence, node *ast.Node) bool {
		s, _ := node.String()
		modArr = append(modArr, s)
		return true
	})

	sort.Sort(sortModArr(modArr))

	mod := ""
	for len(modArr) > 0 {
		randIdx := util.RandInt(0, len(modArr))
		if !isItemIncompatibleWithCurrentItem(items, modTpl, modSlot) {
			mod = modArr[randIdx]
			break
		}

		modArr = append(modArr[:randIdx], modArr[randIdx+1:]...)
	}

	return mod
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
				maxDurability = getRandomisedMaxWeaponOrArmorDurability(botRole, "armor")
				currentDurability = getRandomisedWeaponOrArmorDurability(botRole, "armor", maxDurability)
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
			props.Set("FireMode", ast.NewObject([]ast.Pair{{Key: "FireMode", Value: ast.NewString(util.RandChoose(ft).(string))}}))
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

func generateInventoryBase() *ast.Node {
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

	items := ast.NewArray(nil)
	for _, group := range [][]string{
		{equipID, equipTpl},
		{stashID, stashTpl},
		{questRaidItemID, questRaidItemTpl},
		{questStashItemsID, questStashItemsTpl},
		{sortingTableID, sortingTableTpl},
	} {
		items.Add(
			ast.NewObject([]ast.Pair{
				{
					Key:   "_id",
					Value: ast.NewString(group[0]),
				},
				{
					Key:   "_tpl",
					Value: ast.NewString(group[1]),
				},
			}))
	}

	ivtBase := ast.NewObject(nil)
	ivtBase.Set("items", items)
	ivtBase.Set("equipment", ast.NewString(equipID))
	ivtBase.Set("stash", ast.NewString(stashID))
	ivtBase.Set("questRaidItems", ast.NewString(questRaidItemID))
	ivtBase.Set("questStashItems", ast.NewString(questStashItemsID))
	ivtBase.Set("sortingTable", ast.NewString(sortingTableID))
	ivtBase.Set("fastPanel", ast.NewObject(nil))

	return &ivtBase
}
