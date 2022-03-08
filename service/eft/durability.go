package eft

import (
	"github.com/donkeywon/eft-spg/service/cfg"
	"github.com/donkeywon/eft-spg/util"
)

func getRandomisedMaxWeaponOrArmorDurability(botRole string, weaponOrArmor string) int {
	if isBotPmc(botRole) {
		return generateRandomMaxWeaponOrArmorDurability("pmc", weaponOrArmor)
	} else if isBotBoss(botRole) {
		return generateRandomMaxWeaponOrArmorDurability("boss", weaponOrArmor)
	} else if isBotFollower(botRole) {
		return generateRandomMaxWeaponOrArmorDurability("follower", weaponOrArmor)
	} else {
		return generateRandomMaxWeaponOrArmorDurability(botRole, weaponOrArmor)
	}
}

func generateRandomMaxWeaponOrArmorDurability(botRole string, weaponOrArmor string) int {
	return util.RandInt(getMaxWeaponOrArmorDurabilityFromConfig(botRole, weaponOrArmor, true), getMaxWeaponOrArmorDurabilityFromConfig(botRole, weaponOrArmor, false))
}

func getMaxWeaponOrArmorDurabilityFromConfig(botRole string, weaponOrArmor string, lowest bool) int {
	typ := "lowest"
	if !lowest {
		typ = "highest"
	}

	if cfg.GetCfg().GetByPath("durability", botRole).Exists() {
		l, _ := cfg.GetCfg().GetByPath("bot", "durability", botRole, weaponOrArmor, typ).Int64()
		return int(l)
	}
	l, _ := cfg.GetCfg().GetByPath("bot", "durability", "default", weaponOrArmor, typ).Int64()
	return int(l)
}

func getRandomisedWeaponOrArmorDurability(botRole string, weaponOrArmor string, maxDurability int) int {
	if isBotPmc(botRole) {
		return generateWeaponOrArmorDurability("pmc", weaponOrArmor, maxDurability)
	} else if isBotBoss(botRole) {
		return generateWeaponOrArmorDurability("boss", weaponOrArmor, maxDurability)
	} else if isBotFollower(botRole) {
		return generateWeaponOrArmorDurability("follower", weaponOrArmor, maxDurability)
	} else {
		return generateWeaponOrArmorDurability(botRole, weaponOrArmor, maxDurability)
	}
}

func generateWeaponOrArmorDurability(botRole string, weaponOrArmor string, maxDurability int) int {
	minDelta := getWeaponOrArmorDeltaFromConfig(botRole, true, weaponOrArmor)
	maxDelta := getWeaponOrArmorDeltaFromConfig(botRole, false, weaponOrArmor)

	return maxDurability - util.RandInt(minDelta, maxDelta)
}

func getWeaponOrArmorDeltaFromConfig(botRole string, isMin bool, weaponOrArmor string) int {
	typ := "minDelta"
	if !isMin {
		typ = "maxDelta"
	}

	if cfg.GetCfg().GetByPath("bot", "durability", botRole).Exists() {
		n, _ := cfg.GetCfg().GetByPath("bot", "durability", botRole, weaponOrArmor, typ).Int64()
		return int(n)
	}

	n, _ := cfg.GetCfg().GetByPath("bot", "durability", "default", weaponOrArmor, typ).Int64()
	return int(n)
}
