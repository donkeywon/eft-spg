package controller

import "net/http"

func init() {
	RegisterRouter("/raid/profile/list", GetProfile)
	RegisterRouter("/client/match/available", ServerAvailable)
	RegisterRouter("/client/match/updatePing", UpdatePing)
	RegisterRouter("/client/match/join", JoinMatch)
	RegisterRouter("/client/match/group/start_game", JoinMatch)
	RegisterRouter("/client/match/exit", ExitMatch)
	RegisterRouter("/client/match/group/create", CreateGroup)
	RegisterRouter("/client/match/group/delete", DeleteGroup)
	RegisterRouter("/client/match/group/status", GetGroupStatus)
	RegisterRouter("/client/match/group/exit_from_menu", ExitToMenu)
	RegisterRouter("/client/match/group/looking/start", StartGroupSearch)
	RegisterRouter("/client/match/group/looking/stop", StopGroupSearch)
	RegisterRouter("/client/match/group/invite/send", SendGroupInvite)
	RegisterRouter("/client/match/group/invite/accept", AcceptGroupInvite)
	RegisterRouter("/client/match/group/invite/cancel", CancelGroupInvite)
	RegisterRouter("/client/match/offline/start", StartOfflineRaid)
	RegisterRouter("/client/match/offline/end", EndOfflineRaid)
	RegisterRouter("/client/putMetrics", PutMetrics)
	RegisterRouter("/client/getMetricsConfig", GetMetrics)
}

func GetProfile(resp http.ResponseWriter, req *http.Request) {

}

func ServerAvailable(resp http.ResponseWriter, req *http.Request) {

}

func UpdatePing(resp http.ResponseWriter, req *http.Request) {

}

func JoinMatch(resp http.ResponseWriter, req *http.Request) {

}

func ExitMatch(resp http.ResponseWriter, req *http.Request) {

}

func CreateGroup(resp http.ResponseWriter, req *http.Request) {

}

func DeleteGroup(resp http.ResponseWriter, req *http.Request) {

}

func GetGroupStatus(resp http.ResponseWriter, req *http.Request) {

}

func ExitToMenu(resp http.ResponseWriter, req *http.Request) {

}

func StartGroupSearch(resp http.ResponseWriter, req *http.Request) {

}

func StopGroupSearch(resp http.ResponseWriter, req *http.Request) {

}

func SendGroupInvite(resp http.ResponseWriter, req *http.Request) {

}

func AcceptGroupInvite(resp http.ResponseWriter, req *http.Request) {

}

func CancelGroupInvite(resp http.ResponseWriter, req *http.Request) {

}

func StartOfflineRaid(resp http.ResponseWriter, req *http.Request) {

}

func EndOfflineRaid(resp http.ResponseWriter, req *http.Request) {

}

func PutMetrics(resp http.ResponseWriter, req *http.Request) {

}

func GetMetrics(resp http.ResponseWriter, req *http.Request) {

}
