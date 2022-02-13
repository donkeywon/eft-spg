package httpd

const (
	DefaultIP   = "127.0.0.1"
	DefaultPort = 6969
)

type Config struct {
	IP   string `json:"ip" yaml:"ip"`
	Port int    `json:"port" yaml:"port"`
}

func NewConfig() *Config {
	return &Config{
		IP:   DefaultIP,
		Port: DefaultPort,
	}
}
