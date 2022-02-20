package database

const (
	DefaultPath = "./assets/database"
)

type Config struct {
	Path string `yaml:"path" json:"path"`
}

func NewConfig() *Config {
	return &Config{Path: DefaultPath}
}
