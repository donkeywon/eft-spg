package controller

import "net/http"

func init() {
	RegisterRouter("/client/locations", GetLocationData)
	RegisterRouter("/client/location/getLocalloot", GetLocation)
}

func GetLocationData(resp http.ResponseWriter, req *http.Request) {

}

func GetLocation(resp http.ResponseWriter, req *http.Request) {

}
