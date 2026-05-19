package config

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func SetupLogger(appName, logPath string, debug bool) (*zap.Logger, error) {
	if err := os.MkdirAll(logPath, 0755); err != nil {
		return nil, err
	}

	rotate := &lumberjack.Logger{
		Filename:   filepath.Join(logPath, "app.log"),
		MaxSize:    50,    
		MaxBackups: 5,     
		MaxAge:     30,    
		Compress:   true,
		LocalTime:  true, 
	}

	
	logLevel := zap.NewAtomicLevel()
	if debug {
		logLevel.SetLevel(zap.DebugLevel)
	} else {
		logLevel.SetLevel(zap.InfoLevel)
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.CallerKey = "caller"
	encoderConfig.StacktraceKey = "stacktrace"
	
	var encoder zapcore.Encoder
	if debug {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(rotate), 
			zapcore.AddSync(os.Stdout), 
		),
		logLevel,
	)

	logger := zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.ErrorLevel),
		zap.Fields(
			zap.String("app", appName),
			zap.String("env", map[bool]string{true: "development", false: "production"}[debug]),
		),
	)

	return logger, nil
}