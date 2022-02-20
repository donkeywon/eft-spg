package httpd

import (
	"net/http"
)

func (s *Svc) registerWeatherRouter() {
	s.RegisterRouter("/client/weather", s.GetWeather)
}

func (s *Svc) GetWeather(resp http.ResponseWriter, req *http.Request) {

}
