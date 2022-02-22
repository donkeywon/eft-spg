package httpd

import (
	"github.com/bytedance/sonic/ast"
	"net/http"
)

func (s *Svc) registerWeatherRouter() {
	s.RegisterRouter("/client/weather", s.GetWeather, true)
}

func (s *Svc) GetWeather(sessID string, body *ast.Node, r *http.Request) (interface{}, error) {
	return nil, nil

}
