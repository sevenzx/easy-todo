package setup

import (
	"easytodo/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func Zap() (logger *zap.Logger) {
	lumberjack := Lumberjack()
	level := config.Zap.LogLevel()
	encoder := config.Zap.Encoder()

	lumberjackCore := zapcore.NewCore(encoder, zapcore.AddSync(lumberjack), level)
	cores := make([]zapcore.Core, 0)
	cores = append(cores, lumberjackCore)
	if config.Zap.LogInConsole {
		consoleCore := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level)
		cores = append(cores, consoleCore)
	}
	core := zapcore.NewTee(cores...)
	logger = zap.New(core)
	if config.Zap.AddCaller {
		logger = logger.WithOptions(zap.AddCaller())
	}
	defer logger.Sync()
	return logger
}
