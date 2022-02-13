package controller

import "net/http"

func init() {
	RegisterRouter("/client/ragfair/search", Search)
	RegisterRouter("/client/ragfair/find", Search)
	RegisterRouter("/client/ragfair/itemMarketPrice", GetMarketPrice)
	RegisterRouter("/client/items/prices", GetItemPrices)
	RegisterRouter("/client/trading/api/traderSettings", GetTraderSettings)
}

func Search(resp http.ResponseWriter, req *http.Request) {

}

func GetMarketPrice(resp http.ResponseWriter, req *http.Request) {

}

func GetItemPrices(resp http.ResponseWriter, req *http.Request) {

}

func GetTraderSettings(resp http.ResponseWriter, req *http.Request) {

}
