package controller

import "net/http"

func init() {
	RegisterRouter("/client/trading/api/getUserAssortPrice/trader/{traderID}", GetProfilePurchases)
	RegisterRouter("/client/trading/api/getTrader/{traderID}", GetTrader)
	RegisterRouter("/client/trading/api/getTraderAssort/{traderID}", GetAssort)
}

func GetProfilePurchases(resp http.ResponseWriter, req *http.Request) {

}

func GetTrader(resp http.ResponseWriter, req *http.Request) {

}

func GetAssort(resp http.ResponseWriter, req *http.Request) {

}
