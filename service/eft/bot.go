package eft

import (
	"eft-spg/service/cfg"
	"eft-spg/service/database"
	"eft-spg/util"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"math/rand"
	"strings"
)

const (
	SideSavage  = "Savage"
	SidePmcUsec = "Usec"
	SidePmcBear = "Bear"
)

func GetBotLimit(typ string) (int64, error) {
	if typ == "" {
		return 0, errors.New(util.ErrIllegalArg)
	}

	if typ == "cursedAssault" || typ == "assaultGroup" {
		typ = "assault"
	}

	l := cfg.GetCfg().GetByPath("bot", "presetBatch", typ)
	if l == nil {
		return 0, errors.New(util.ErrIllegalArg)
	}

	return l.Int64()
}

func GetBotDifficulty(typ string, difficulty string) (*ast.Node, error) {
	bearType, _ := cfg.GetCfg().GetByPath("bot", "pmc", "bearType").String()
	usecType, _ := cfg.GetCfg().GetByPath("bot", "pmc", "usecType").String()
	chanceSameSideIsHostilePercent, _ := cfg.GetCfg().GetByPath("bot", "pmc", "chanceSameSideIsHostilePercent").Int64()

	switch typ {
	case "core":
		return database.GetDatabase().GetByPath("bots", "core"), nil
	case bearType, usecType:
		difficultySettings := GetPmcDifficultySettings(typ, difficulty)
		if rand.Int63n(100) < chanceSameSideIsHostilePercent {
			difficultySettings.Get("Mind").SetAny("DEFAULT_ENEMY_USEC", true)
			difficultySettings.Get("Mind").SetAny("DEFAULT_ENEMY_BEAR", true)
		}
		return difficultySettings, nil
	default:
		return database.GetDatabase().GetByPath("bots", "types", typ, "difficulty", difficulty), nil
	}
}

func GetPmcDifficultySettings(typ string, difficulty string) *ast.Node {
	pmcD, _ := cfg.GetCfg().GetByPath("bot", "pmc", "difficulty").String()
	if strings.ToLower(pmcD) != "asonline" {
		difficulty = pmcD
	}

	return database.GetDatabase().GetByPath("bots", "types", typ, "difficulty", difficulty)
}

func GetBotCap() int64 {
	c, _ := cfg.GetCfg().GetByPath("bot", "maxBotCap").Int64()
	return c
}

func GetPmcDifficulty(d string) string {
	bd, _ := cfg.GetCfg().GetByPath("bot", "pmc", "difficulty").String()
	if strings.ToLower(bd) == "asonline" {
		return d
	}
	return bd
}

func Generate(info *ast.Node, playerScav bool) *ast.Node {
	out := ast.NewArray(nil)
	if !info.Get("conditions").Exists() {
		return &out
	}

	isUsecChance, _ := cfg.GetCfg().GetByPath("pmc", "isUsec").Int64()

	conditions, _ := info.Get("conditions").ArrayUseNode()
	count := 0
	for _, condition := range conditions {
		if !condition.Get("Limit").Exists() {
			continue
		}
		limit, _ := condition.Get("Limit").Int64()

		for i := 0; i < int(limit); i++ {
			side := SidePmcUsec
			if rand.Int63n(100) < isUsecChance {
				side = "Bear"
			}
			role, _ := condition.Get("Role").String()
			isPmc := false
			if !playerScav {
				if cfg.GetCfg().GetByPath("bot", "pmc", "types", role).Exists() {
					n, _ := cfg.GetCfg().GetByPath("bot", "pmc", "types", role).Int64()
					if rand.Int63n(100) < n {
						isPmc = true
					}
				}
			}

			cd, _ := condition.Get("Difficulty").String()
			bot := database.GetDatabase().GetByPath("bots", "base")
			if isPmc {
				bot.GetByPath("Info", "Settings").SetAny("BotDifficulty", GetPmcDifficulty(cd))

				if side == SidePmcUsec {
					role, _ = cfg.GetCfg().GetByPath("bot", "pmc", "usecType").String()
				} else if side == SidePmcBear {
					role, _ = cfg.GetCfg().GetByPath("bot", "pmc", "bearType").String()
				}
			} else {
				bot.GetByPath("Info", "Settings").SetAny("BotDifficulty", cd)
			}

			bot.GetByPath("Info", "Settings").SetAny("Role", role)
			if isPmc {
				bot.Get("Info").SetAny("Side", side)
			} else {
				bot.Get("Info").SetAny("Side", SideSavage)
			}

			// TODO generateBot

			out.Add(*bot)
			count++
		}
	}

	GetSvc().Info("Generate bot", zap.Int("num", count))
	return &out
}

func generateBot(bot *ast.Node, role string, isPmc bool) {
	roleBot := database.GetDatabase().GetByPath("bots", "types", strings.ToLower(role))
	minLvl, _ := roleBot.GetByPath("experience", "level", "min").Int64()
	maxLvl, _ := roleBot.GetByPath("experience", "level", "max").Int64()

	firstNames, _ := roleBot.Get("firstName").Array()
	lastNames, _ := roleBot.Get("lastName").Array()

	firstName := util.RandChoose(firstNames)
	lastName := util.RandChoose(lastNames)
	if lastName == nil {
		lastName = ""
	}

	name := firstName.(string) + " " + lastName.(string)
	showTypeInNickname, _ := cfg.GetCfg().GetByPath("bot", "showTypeInNickName").Bool()
	if showTypeInNickname {
		name += " " + role
	}

	bot.Get("Info").SetAny("Nickname", name)

	if !ChristmasEventEnabled() {
		for _, n := range []*ast.Node{roleBot.GetByPath("Inventory", "equipment"), roleBot.GetByPath("Inventory", "items")} {
			n.ForEach(func(path ast.Sequence, node *ast.Node) bool {
				m, _ := node.Map()
				for k, _ := range m {
					if ItemIsChristmasRelated(k) {
						delete(m, k)
					}
				}

				n.SetAny(*path.Key, m)

				return true
			})
		}
	}

	lvl, exp := generateRandomLevel(int(minLvl), int(maxLvl))
	minRewardExp, _ := roleBot.GetByPath("experience", "reward", "min").Int64()
	maxRewardExp, _ := roleBot.GetByPath("experience", "reward", "max").Int64()
	voices, _ := roleBot.GetByPath("appearance", "voice").Array()
	side, _ := bot.GetByPath("Info", "Side").String()

	bot.Get("Info").SetAny("Experience", exp)
	bot.Get("Info").SetAny("Level", lvl)
	bot.GetByPath("Info", "Settings").SetAny("Experience", util.RandInt(int(minRewardExp), int(maxRewardExp)))
	bot.GetByPath("Info", "Settings").Set("StandingForKill", *roleBot.GetByPath("experience", "standingForKill"))
	bot.GetByPath("Info").SetAny("Voice", util.RandChoose(voices))
	bot.Set("Health", generateHealth(roleBot.Get("health"), side == SideSavage))
	bot.Set("Skills", generateSkills(roleBot.Get("skills")))
	for _, t := range []string{"Head", "Body", "Feet", "Hands"} {
		bot.Get("Customization").Set(t, *util.RandChooseNode(roleBot.GetByPath("appearance", strings.ToLower(t))))
	}
	// TODO generateInventory
}

func generateRandomLevel(min int, max int) (int, int) {
	expN, _ := database.GetDatabase().GetByPath("globals", "config", "exp", "level", "exp_table").ArrayUseNode()
	maxLvl := max
	if len(expN) < max {
		maxLvl = len(expN)
	}

	exp := 0
	lvl := util.RandInt(min, maxLvl)

	for i := 0; i < lvl; i++ {
		e, _ := expN[i].Get("exp").Int64()
		exp += int(e)
	}

	if lvl < len(expN)-1 {
		e, _ := expN[lvl].Get("exp").Int64()
		exp += util.RandInt(0, int(e)-1)
	}

	return lvl, exp
}

func generateHealth(health *ast.Node, playerScav bool) ast.Node {
	bodyParts := health.GetByPath("BodyParts", 0)
	if !playerScav {
		bodyParts = util.RandChooseNode(health.Get("BodyParts"))
	}

	n, _ := sonic.Get([]byte(`{
"Hydration": {
    "Current": 0,
    "Maximum": 0
},
"Energy": {
    "Current": 0,
    "Maximum": 0
},
"Temperature": {
    "Current": 0,
    "Maximum": 0
},
"BodyParts": {
    "Head": {
        "Health": {
            "Current": 0,
            "Maximum": 0
        }
    },
    "Chest": {
        "Health": {
            "Current": 0,
            "Maximum": 0
        }
    },
    "Stomach": {
        "Health": {
            "Current": 0,
            "Maximum": 0
        }
    },
    "LeftArm": {
        "Health": {
            "Current": 0,
            "Maximum": 0
        }
    },
    "RightArm": {
        "Health": {
            "Current": 0,
            "Maximum": 0
        }
    },
    "LeftLeg": {
        "Health": {
            "Current": 0,
            "Maximum": 0
        }
    },
    "RightLeg": {
        "Health": {
            "Current": 0,
            "Maximum": 0
        }
    }
}
}`))

	for _, t := range []string{"Hydration", "Energy", "Temperature"} {
		n.Get(t).SetAny("Current", util.RandIntNode(health.GetByPath(t, "min"), health.GetByPath(t, "max")))
		n.Get(t).Set("Maximum", *health.GetByPath(t, "max"))
	}

	bp := n.Get("BodyParts")
	for _, t := range []string{"Head", "Chest", "Stomach", "LeftArm", "RightArm", "LeftLeg", "RightLeg"} {
		bp.GetByPath(t, "Health").SetAny("Current", util.RandIntNode(bodyParts.GetByPath(t, "min"), bodyParts.GetByPath(t, "max")))
		bp.GetByPath(t, "Health").Set("Maximum", *bodyParts.GetByPath("Head", "max"))
	}

	return n
}

func generateSkills(skillsNode *ast.Node) ast.Node {
	skills := ast.NewArray(nil)
	masteries := ast.NewArray(nil)

	for _, typ := range []string{"Common", "Mastering"} {
		if skillsNode.Get(typ).Exists() {
			skillsNode.Get(typ).ForEach(func(path ast.Sequence, node *ast.Node) bool {
				skills.Add(ast.NewObject([]ast.Pair{{
					Key:   "Id",
					Value: ast.NewString(*path.Key),
				}, {
					Key:   "Progress",
					Value: ast.NewAny(util.RandIntNode(skillsNode.GetByPath(typ, path.Key, "min"), skillsNode.GetByPath(typ, path.Key, "max"))),
				}}))

				return true
			})
		}
	}

	return ast.NewObject([]ast.Pair{{
		Key:   "Common",
		Value: skills,
	}, {
		Key:   "Mastering",
		Value: masteries,
	}, {
		Key:   "Points",
		Value: ast.NewAny(0),
	}})
}
