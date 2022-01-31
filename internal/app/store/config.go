package store

// Config ...
type Config struct {
	DatabaseURl string `toml:"database_url"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
