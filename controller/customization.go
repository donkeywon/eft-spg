package controller

import "net/http"

func init() {
	RegisterRouter("/client/trading/customization/storage", GetSuits)
	RegisterRouter("/client/trading/customization/{traderID}/{aaa}", GetTraderSuits)
}

func GetSuits(resp http.ResponseWriter, req *http.Request) {

}

func GetTraderSuits(resp http.ResponseWriter, req *http.Request) {

}
