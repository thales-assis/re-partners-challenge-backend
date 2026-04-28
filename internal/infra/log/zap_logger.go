package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ProvideZapLogger() (*zap.Logger, error) {

	opts := make([]zap.Option, 0)

	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	logger, err := config.Build(opts...)
	if err != nil {
		return nil, err
	}

	return logger, nil
}
