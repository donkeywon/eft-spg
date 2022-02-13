package controller

import "net/http"

func init() {
	RegisterRouter("/client/weather", GetWeather)
}

func GetWeather(resp http.ResponseWriter, req *http.Request) {

}
