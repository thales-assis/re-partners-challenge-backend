package config

type Config struct {
	Server        ServerConfig        `mapstructure:"server"`
	SystemManager SystemManagerConfig `mapstructure:"system-manager"`
}

func ProvideConfig(
	c ConfigReader,
) (*Config, error) {

	var config Config

	if err := c.ReadDirectlyFromFile(&config); err != nil {
		return nil, err
	}

	if config.SystemManager.Enabled {
		// Implement the flow to get the configuration from the System Manager (secrets parameters, secrets configurations, etc)
		// For example -> err := c.LoadFromSSMParameterStore(&config)
	}

	return &config, nil
}
