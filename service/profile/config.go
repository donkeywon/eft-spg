package profile

const (
	DefaultPath = "user/profiles/"
)

type Config struct {
	Path string `yaml:"path" json:"path"`
}

func NewConfig() *Config {
	return &Config{Path: DefaultPath}
}
