package apiserver

import "github.com/Tradash/gopherSchool/internal/app/store"

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"logs_level"`
	Store    *store.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{BindAddr: ":8080", LogLevel: "debug", Store: store.NewConfig()}

}
