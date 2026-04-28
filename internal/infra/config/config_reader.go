package config

import (
	"github.com/re-partners-challenge-backend/internal/infra/log"
	"github.com/spf13/viper"
)

const (
	ViperConfigurationName    = "config"
	ViperDefaultFilePath      = "."
	ViperDefaultFileExtension = "toml"
)

type ConfigReader struct {
	logger *log.ZapLogger
}

func ProvideConfigReader(
	logger *log.ZapLogger,
) ConfigReader {
	return ConfigReader{
		logger,
	}
}

func (c ConfigReader) ReadDirectlyFromFile(config *Config) error {

	v := viper.New()
	v.SetConfigName(ViperConfigurationName)
	v.AddConfigPath(ViperDefaultFilePath)
	v.AutomaticEnv()
	v.SetConfigType(ViperDefaultFileExtension)

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.Unmarshal(config); err != nil {
		return err
	}

	return nil
}
