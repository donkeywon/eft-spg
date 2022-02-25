package eft

import (
	"eft-spg/service/cfg"
	"eft-spg/service/database"
	"eft-spg/util"
	"github.com/bytedance/sonic/ast"
	"github.com/pkg/errors"
	"math/rand"
	"strings"
)

func (s *Svc) GetBotLimit(typ string) (int64, error) {
	if typ == "" {
		return 0, errors.New(util.ErrIllegalArg)
	}

	if typ == "cursedAssault" || typ == "assaultGroup" {
		typ = "assault"
	}

	l := cfg.GetSvc().GetConfig().GetByPath("bot", "presetBatch", typ)
	if l == nil {
		return 0, errors.New(util.ErrIllegalArg)
	}

	return l.Int64()
}

func (s *Svc) GetBotDifficulty(typ string, difficulty string) (*ast.Node, error) {
	bearType, _ := cfg.GetSvc().GetConfig().GetByPath("bot", "pmc", "bearType").String()
	usecType, _ := cfg.GetSvc().GetConfig().GetByPath("bot", "pmc", "usecType").String()
	chanceSameSideIsHostilePercent, _ := cfg.GetSvc().GetConfig().GetByPath("bot", "pmc", "chanceSameSideIsHostilePercent").Int64()

	switch typ {
	case "core":
		return database.GetSvc().GetDatabase().GetByPath("bots", "core"), nil
	case bearType, usecType:
		difficultySettings := s.GetPmcDifficultySettings(typ, difficulty)
		if rand.Int63n(100) < chanceSameSideIsHostilePercent {
			difficultySettings.Get("Mind").SetAny("DEFAULT_ENEMY_USEC", true)
			difficultySettings.Get("Mind").SetAny("DEFAULT_ENEMY_BEAR", true)
		}
		return difficultySettings, nil
	default:
		return database.GetSvc().GetDatabase().GetByPath("bots", "types", typ, "difficulty", difficulty), nil
	}
}

func (s *Svc) GetPmcDifficultySettings(typ string, difficulty string) *ast.Node {
	pmcD, _ := cfg.GetSvc().GetConfig().GetByPath("bot", "pmc", "difficulty").String()
	if strings.ToLower(pmcD) != "asonline" {
		difficulty = pmcD
	}

	return database.GetSvc().GetDatabase().GetByPath("bots", "types", typ, "difficulty", difficulty)
}

func (s *Svc) GetBotCap() int64 {
	c, _ := cfg.GetSvc().GetConfig().GetByPath("bot", "maxBotCap").Int64()
	return c
}
