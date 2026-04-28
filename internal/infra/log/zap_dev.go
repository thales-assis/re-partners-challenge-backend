//go:build !prod
// +build !prod

package log

import (
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.Logger
}

func ProvideLogger() (*ZapLogger, error) {
	zap, _ := zap.NewDevelopment()
	return &ZapLogger{logger: zap}, nil
}

func (z *ZapLogger) Debug(msg string, fields ...LoggerField) {
	z.logger.Debug(msg, z.parseToZapField(fields)...)
}

func (z *ZapLogger) Error(msg string, fields ...LoggerField) {
	z.logger.Error(msg, z.parseToZapField(fields)...)
}

func (z *ZapLogger) Fatal(msg string, fields ...LoggerField) {
	z.logger.Fatal(msg, z.parseToZapField(fields)...)
}

func (z *ZapLogger) Info(msg string, fields ...LoggerField) {
	z.logger.Info(msg, z.parseToZapField(fields)...)
}

func (z *ZapLogger) Warn(msg string, fields ...LoggerField) {
	z.logger.Warn(msg, z.parseToZapField(fields)...)
}

func (z *ZapLogger) With(fields ...LoggerField) *ZapLogger {
	return nil
}

func (z *ZapLogger) parseToZapField(fields []LoggerField) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		value, ok := f.FieldValue.(error)
		if ok {
			zapFields = append(zapFields, zap.Error(value))
			continue
		}
		zapFields = append(zapFields, zap.Any(f.FieldName, f.FieldValue))
	}
	return zapFields
}
