package hook

import (
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/donkeywon/gtil/util"
)

func init() {
	RegisterPostLoadHook(resetVitality)
}

func resetVitality(profile *ast.Node) {
	if profile == nil {
		return
	}

	str := `
{
    "health": {
        "Hydration": 0,
        "Energy": 0,
        "Temperature": 0,
        "Head": 0,
        "Chest": 0,
        "Stomach": 0,
        "LeftArm": 0,
        "RightArm": 0,
        "LeftLeg": 0,
        "RightLeg": 0
    },
    "effects": {
        "Head": {},
        "Chest": {},
        "Stomach": {},
        "LeftArm": {},
        "RightArm": {},
        "LeftLeg": {},
        "RightLeg": {}
    }
}`
	n, _ := sonic.Get(util.String2Bytes(str))

	profile.Set("vitality", n)
}
