package config

type Config struct {
	Server ServerConfig `mapstructure:"server"`
}

func ProvideConfig(
	c ConfigReader,
) (*Config, error) {

	var config Config

	if err := c.ReadDirectlyFromFile(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
