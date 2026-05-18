package helpers

import "go.uber.org/zap"


var globalLogger *zap.Logger

func SetGlobalLogger(logger *zap.Logger) {
	globalLogger = logger
}

func Info(msg string, fields ...zap.Field) {
	if globalLogger != nil {
		globalLogger.Info(msg, fields...)
	}
}

func Error(msg string, fields ...zap.Field) {
	if globalLogger != nil {
		globalLogger.Error(msg, fields...)
	}
}

func Debug(msg string, fields ...zap.Field) {
	if globalLogger != nil {
		globalLogger.Debug(msg, fields...)
	}
}

func Warn(msg string, fields ...zap.Field) {
	if globalLogger != nil {
		globalLogger.Warn(msg, fields...)
	}
}

func WithContext(fields ...zap.Field) *zap.Logger {
	if globalLogger != nil {
		return globalLogger.With(fields...)
	}
	return zap.NewNop()
}