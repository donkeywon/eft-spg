package httpd

import (
	. "eft-spg/controller"
	"github.com/gorilla/mux"
)

func StaticRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/client/game/bot/generate", GenerateBots).Methods("GET", "POST")
	router.HandleFunc("/client/trading/customization/storage", GetSuits).Methods("GET", "POST")
	router.HandleFunc("/client/settings", GetSettings).Methods("GET", "POST")
	router.HandleFunc("/client/globals", GetGlobals).Methods("GET", "POST")
	router.HandleFunc("/client/items", GetTemplateItems).Methods("GET", "POST")
	router.HandleFunc("/client/handbook/templates", GetTemplateHandbook).Methods("GET", "POST")
	router.HandleFunc("/client/customization", GetTemplateSuits).Methods("GET", "POST")
	router.HandleFunc("/client/account/customization", GetTemplateCharacter).Methods("GET", "POST")
	router.HandleFunc("/client/hideout/production/recipes", GetHideoutProduction).Methods("GET", "POST")
	router.HandleFunc("/client/hideout/settings", GetHideoutSettings).Methods("GET", "POST")
	router.HandleFunc("/client/hideout/areas", GetHideoutAreas).Methods("GET", "POST")
	router.HandleFunc("/client/hideout/production/scavcase/recipes", GetHideoutScavcase).Methods("GET", "POST")
	router.HandleFunc("/client/languages", GetLocalesLanguages).Methods("GET", "POST")
	router.HandleFunc("/client/friend/list", GetFriendList).Methods("GET", "POST")
	router.HandleFunc("/client/chatServer/list", GetChatServerList).Methods("GET", "POST")
	router.HandleFunc("/client/mail/dialog/list", GetMailDialogList).Methods("GET", "POST")
	router.HandleFunc("/client/mail/dialog/view", GetMailDialogView).Methods("GET", "POST")
	router.HandleFunc("/client/mail/dialog/info", GetMailDialogInfo).Methods("GET", "POST")
	router.HandleFunc("/client/mail/dialog/remove", RemoveDialog).Methods("GET", "POST")
	router.HandleFunc("/client/mail/dialog/pin", PinDialog).Methods("GET", "POST")
	router.HandleFunc("/client/mail/dialog/unpin", UnpinDialog).Methods("GET", "POST")
	router.HandleFunc("/client/mail/dialog/read", SetRead).Methods("GET", "POST")
	router.HandleFunc("/client/mail/dialog/getAllAttachments", GetAllAttachments).Methods("GET", "POST")
	router.HandleFunc("/client/friend/request/list/outbox", ListOutbox).Methods("GET", "POST")
	router.HandleFunc("/client/friend/request/list/inbox", ListInbox).Methods("GET", "POST")
	router.HandleFunc("/client/friend/request/send", FriendRequest).Methods("GET", "POST")
	router.HandleFunc("/client/game/config", GetGameConfig).Methods("GET", "POST")
	router.HandleFunc("/client/server/list", GetServer).Methods("GET", "POST")
	router.HandleFunc("/client/game/version/validate", VersionValidate).Methods("GET", "POST")
	router.HandleFunc("/client/game/start", GameStart).Methods("GET", "POST")
	router.HandleFunc("/client/game/logout", GameLogout).Methods("GET", "POST")
	router.HandleFunc("/client/checkVersion", ValidateGameVersion).Methods("GET", "POST")
	router.HandleFunc("/client/game/keepalive", GameKeepAlive).Methods("GET", "POST")
	router.HandleFunc("/player/health/sync", SyncHealth).Methods("GET", "POST")
	router.HandleFunc("/raid/profile/save", SaveProgress).Methods("GET", "POST")
	router.HandleFunc("/singleplayer/settings/raid/endstate", GetRaidEndState).Methods("GET", "POST")
	router.HandleFunc("/singleplayer/settings/weapon/durability", GetWeaponDurability).Methods("GET", "POST")
	router.HandleFunc("/singleplayer/settings/raid/menu", GetRaidMenuSettings).Methods("GET", "POST")
	router.HandleFunc("/singleplayer/settings/version", GetVersion).Methods("GET", "POST")
	router.HandleFunc("/client/insurance/items/list/cost", GetInsuranceCost).Methods("GET", "POST")
	router.HandleFunc("/client/game/profile/items/moving", HandleEvents).Methods("GET", "POST")
	router.HandleFunc("/launcher/server/connect", Connect).Methods("GET", "POST")
	router.HandleFunc("/launcher/profile/login", Login).Methods("GET", "POST")
	router.HandleFunc("/launcher/profile/register", Register).Methods("GET", "POST")
	router.HandleFunc("/launcher/profile/get", Get).Methods("GET", "POST")
	router.HandleFunc("/launcher/profile/change/username", ChangeUsername).Methods("GET", "POST")
	router.HandleFunc("/launcher/profile/change/password", ChangePassword).Methods("GET", "POST")
	router.HandleFunc("/launcher/profile/change/wipe", Wipe).Methods("GET", "POST")
	router.HandleFunc("/launcher/profile/info", GetMiniProfile).Methods("GET", "POST")
	router.HandleFunc("/launcher/ping", Ping).Methods("GET", "POST")
	router.HandleFunc("/client/locations", GetLocationData).Methods("GET", "POST")
	router.HandleFunc("/raid/profile/list", GetProfile).Methods("GET", "POST")
	router.HandleFunc("/client/match/available", ServerAvailable).Methods("GET", "POST")
	router.HandleFunc("/client/match/updatePing", UpdatePing).Methods("GET", "POST")
	router.HandleFunc("/client/match/join", JoinMatch).Methods("GET", "POST")
	router.HandleFunc("/client/match/exit", ExitMatch).Methods("GET", "POST")
	router.HandleFunc("/client/match/group/create", CreateGroup).Methods("GET", "POST")
	router.HandleFunc("/client/match/group/delete", DeleteGroup).Methods("GET", "POST")
	router.HandleFunc("/client/match/group/status", GetGroupStatus).Methods("GET", "POST")
	router.HandleFunc("/client/match/group/start_game", JoinMatch).Methods("GET", "POST")
	router.HandleFunc("/client/match/group/exit_from_menu", ExitToMenu).Methods("GET", "POST")
	router.HandleFunc("/client/match/group/looking/start", StartGroupSearch).Methods("GET", "POST")
	router.HandleFunc("/client/match/group/looking/stop", StopGroupSearch).Methods("GET", "POST")
	router.HandleFunc("/client/match/group/invite/send", SendGroupInvite).Methods("GET", "POST")
	router.HandleFunc("/client/match/group/invite/accept", AcceptGroupInvite).Methods("GET", "POST")
	router.HandleFunc("/client/match/group/invite/cancel", CancelGroupInvite).Methods("GET", "POST")
	router.HandleFunc("/client/match/offline/start", StartOfflineRaid).Methods("GET", "POST")
	router.HandleFunc("/client/match/offline/end", EndOfflineRaid).Methods("GET", "POST")
	router.HandleFunc("/client/putMetrics", PutMetrics).Methods("GET", "POST")
	router.HandleFunc("/client/getMetricsConfig", GetMetrics).Methods("GET", "POST")
	router.HandleFunc("/singleplayer/bundles", GetBundles).Methods("GET", "POST")
	router.HandleFunc("/client/notifier/channel/create", CreateNotifierChannel).Methods("GET", "POST")
	router.HandleFunc("/client/game/profile/select", SelectProfile).Methods("GET", "POST")
	router.HandleFunc("/client/handbook/builds/my/list", GetHandbookUserList).Methods("GET", "POST")
	router.HandleFunc("/client/game/profile/create", CreateProfile).Methods("GET", "POST")
	router.HandleFunc("/client/game/profile/list", GetProfileData).Methods("GET", "POST")
	router.HandleFunc("/client/game/profile/savage/regenerate", ReGenerateScav).Methods("GET", "POST")
	router.HandleFunc("/client/game/profile/voice/change", ChangeVoice).Methods("GET", "POST")
	router.HandleFunc("/client/game/profile/nickname/change", ChangeNickname).Methods("GET", "POST")
	router.HandleFunc("/client/game/profile/nickname/validate", ValidateNickname).Methods("GET", "POST")
	router.HandleFunc("/client/game/profile/nickname/reserved", GetReservedNickname).Methods("GET", "POST")
	router.HandleFunc("/client/profile/status", GetProfileStatus).Methods("GET", "POST")
	router.HandleFunc("/client/quest/list", ListQuests).Methods("GET", "POST")
	router.HandleFunc("/client/repeatalbeQuests/activityPeriods", ActivityPeriods).Methods("GET", "POST")
	router.HandleFunc("/client/ragfair/search", Search).Methods("GET", "POST")
	router.HandleFunc("/client/ragfair/find", Search).Methods("GET", "POST")
	router.HandleFunc("/client/ragfair/itemMarketPrice", GetMarketPrice).Methods("GET", "POST")
	router.HandleFunc("/client/items/prices", GetItemPrices).Methods("GET", "POST")
	router.HandleFunc("/client/trading/api/traderSettings", GetTraderSettings).Methods("GET", "POST")
	router.HandleFunc("/client/weather", GetWeather).Methods("GET", "POST")

	router.HandleFunc("/client/menu/locale/{locale}", GetLocalesMenu).Methods("GET", "POST")
	router.HandleFunc("/client/locale/{locale}", GetLocalesGlobal).Methods("GET", "POST")
	router.HandleFunc("/singleplayer/settings/bot/limit/{type}", GetBotLimit).Methods("GET", "POST")
	router.HandleFunc("/singleplayer/settings/bot/difficulty/{type}/{difficulty}", GetBotDifficulty).Methods("GET", "POST")
	router.HandleFunc("/client/trading/customization/{traderID}/{aaa}", GetTraderSuits).Methods("GET", "POST")
	router.HandleFunc(".jpg", GetImage).Methods("GET", "POST")
	router.HandleFunc(".png", GetImage).Methods("GET", "POST")
	router.HandleFunc(".ico", GetImage).Methods("GET", "POST")
	router.HandleFunc("/client/location/getLocalloot", GetLocation).Methods("GET", "POST")
	router.HandleFunc(".bundle", GetBundles).Methods("GET", "POST")
	router.HandleFunc("/client/trading/api/getUserAssortPrice/trader/{traderID}", GetProfilePurchases).Methods("GET", "POST")
	router.HandleFunc("/client/trading/api/getTrader/{traderID}", GetTrader).Methods("GET", "POST")
	router.HandleFunc("/client/trading/api/getTraderAssort/{traderID}", GetAssort).Methods("GET", "POST")
	router.HandleFunc("/?last_id", Notify).Methods("GET", "POST")
	router.HandleFunc("/notifierServer", Notify).Methods("GET", "POST")
	router.HandleFunc("/push/notifier/get/", GetNotifier).Methods("GET", "POST")
	router.HandleFunc("/push/notifier/getwebsocket/", GetNotifier).Methods("GET", "POST")
	router.HandleFunc("/singleplayer/settings/bot/maxCap", GetBotCap).Methods("GET", "POST")
	return router
}
