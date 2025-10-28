package api

import "dev-clash/internal/adapters/postgres"

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LoggerLevel string `toml:"logger_level"`
	PostgresURI *postgres.Config `toml:"storage"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LoggerLevel: "debug",
		PostgresURI: &postgres.Config{},
	}
}