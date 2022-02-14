package config

func init() {

}

var c = []byte(`
{
    "presetBatch": {
        "assault": 120,
        "bossBully": 1,
        "bossGluhar": 1,
        "bossKilla": 1,
        "bossKojaniy": 1,
        "bossSanitar": 1,
        "bossTagilla": 1,
        "bossTest": 40,
        "cursedAssault": 120,
        "followerBully": 4,
        "followerGluharAssault": 2,
        "followerGluharScout": 2,
        "followerGluharSecurity": 2,
        "followerGluharSnipe": 2,
        "followerKojaniy": 2,
        "followerSanitar": 2,
        "followerTagilla": 2,
        "followerTest": 4,
        "marksman": 30,
        "pmcBot": 120,
        "sectantPriest": 1,
        "sectantWarrior": 5,
        "gifter": 1,
        "test": 40,
        "exUsec": 15
    },
    "bosses": ["bossbully", "bossgluhar", "bosskilla", "bosskojaniy", "bosssanitar", "bosstagilla"],
    "durability":{
        "pmcbot":{
            "armor": {
                "minPercent": 80
            },
            "weapon": {
                "minPercent": 80
            }
        },
        "exusec":{
            "armor": {
                "minPercent": 80
            },
            "weapon": {
                "minPercent": 80
            }
        },
        "pmc": {
            "armor": {
                "minPercent": 80
            },
            "weapon": {
                "minPercent": 80
            }
        },
        "boss": {
            "armor": {
                "minPercent": 80
            },
            "weapon": {
                "minPercent": 80
            }
        }
    },
    "pmc": {
        "dynamicLoot": {
            "whitelist": [

            ],
            "blacklist": [
                "5fca13ca637ee0341a484f46", // SJ9 TGLabs combat stimulant injector (Thermal Stim)
                "59f32c3b86f77472a31742f0", // usec dogtag
                "59f32bb586f774757e1e8442", // bear dogtag
                "6087e570b998180e9f76dc24" // Superfors DB 2020 Dead Blow Hammer
            ],
            "spawnLimits": {
                "5c99f98d86f7745c314214b3": 1, // mechanical key
                "5c164d2286f774194c5e69fa": 1, // keycard
                "550aa4cd4bdc2dd8348b456c": 2, // silencer
                "55818add4bdc2d5b648b456f": 1, // assault scope
                "55818ad54bdc2ddc698b4569": 1, // reflex sight
                "55818af64bdc2d5b648b4570": 1, // foregrip
                "5448e54d4bdc2dcc718b4568": 1, // armor
                "5448f3a64bdc2d60728b456a": 2, // stims
                "5447e1d04bdc2dff2f8b4567": 1, // knife
                "5a341c4686f77469e155819e": 1, // face cover
                "55818b164bdc2ddc698b456c": 2, // tactical laser/light
                "5448bc234bdc2d3c308b4569": 2, // Magazine
                "543be5dd4bdc2deb348b4569": 2, // Money
                "543be5cb4bdc2deb348b4568": 2 // AmmoBox
            },
            "moneyStackLimits": {
                "5449016a4bdc2d6f028b456f": 4000, // Rouble
                "5696686a4bdc2da3298b456a": 50, // USD
                "569668774bdc2da2298b4568": 50, // Euro
            }
        },
        "cartridgeBlacklist": [
            "56dff421d2720b5f5a8b4567", // 5.45x39mm sp
            "56dff216d2720bbd668b4568", // 5.45x39mm hp
            "56dff338d2720bbd668b4569", // 5.45x39mm prs
            "56dff4ecd2720b5f5a8b4568", // 5.45x39mm US

            "59e6918f86f7746c9f75e849", // 5.56x45mm hp
            "5c0d5ae286f7741e46554302", // 5.56x45mm warmageddon

            "5c0d56a986f774449d5de529", // 9x19mm rip
            "5efb0e16aeb21837e749c7ff", // 9x19mm quakemaker

            "5737218f245977612125ba51", // 9x18mm sp8
            "57372140245977611f70ee91", // 9x18mm sp7
            "57371aab2459775a77142f22", // 9x18mm pmm pstm
            "573719762459775a626ccbc1", // 9x18mm pmp

            "573601b42459776410737435", // 7.62x25mm lrn
            "573602322459776445391df1", // 7.62x25mm lrnpc

            "59e4d3d286f774176a36250a", // 7.62x39mm HP

            "5e023e88277cce2b522ff2b1", // 7.62x51mm ultra nosler

            "59e6658b86f77411d949b250", // .366 tkm

            "5c0d591486f7744c505b416f", // 12/70 rip
            "5d6e68d1a4b93622fe60e845", // 12/70 SuperFormance HP slug
            "5d6e6869a4b9361c140bcfde", // 12/70 Grizzly 40 slug

            "5e85a9f4add9fe03027d9bf1", // 23x75mm flashbang round

            "5cadf6e5ae921500113bb973", // 12.7x55mm PS12A
            "5cadf6ddae9215051e1c23b2", // 12.7x55mm PS12

            "6196365d58ef8c428c287da1", // .300 Blackout Whisper

            "5ba26812d4351e003201fef1", // 4.6x30mm action sx

            "5cc80f79e4a949033c7343b2" // 5.7x28mm SS198LF
        ],
        "difficulty": "AsOnline",
        "isUsec": 50,
        "chanceSameSideIsHostilePercent": 50,
        "usecType": "bosstest",
        "bearType": "test",
        "maxLootTotalRub": 150000,
        "types": {
            "assault": 35,
            "cursedAssault": 35,
            "pmcBot": 35,
            "exUsec": 10
        }
    },
    "showTypeInNickname": false,
    "maxBotCap": 20
}
`)

type bot struct {
}
