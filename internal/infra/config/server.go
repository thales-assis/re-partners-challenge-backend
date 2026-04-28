package config

import "fmt"

type ServerConfig struct {
	Port    int        `mapstructure:"port"`
	Prefix  string     `mapstructure:"prefix"`
	Version string     `mapstructure:"version"`
	Cors    CorsConfig `mapstructure:"cors"`
}

func (cfg ServerConfig) Address() string {
	return fmt.Sprintf(":%d", cfg.Port)
}
