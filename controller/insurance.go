package controller

import "net/http"

func init() {
	RegisterRouter("/client/insurance/items/list/cost", ListInsuranceCost)
}

func ListInsuranceCost(resp http.ResponseWriter, req *http.Request) {

}
