package config

type CorsConfig struct {
	AllowedOrigins []string `mapstructure:"allowed-origins"`
	AllowedMethods []string `mapstructure:"allowed-methods"`
	AllowedHeaders []string `mapstructure:"allowed-headers"`
}
