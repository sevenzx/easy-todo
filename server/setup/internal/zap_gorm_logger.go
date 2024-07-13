package internal

import (
	"context"
	"easytodo/config"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

type ZapGormLogger struct {
	logger *zap.Logger
	config logger.Config
}

func NewZapGormLogger(zapLogger *zap.Logger, lcf logger.Config) *ZapGormLogger {
	if config.Zap.AddCaller {
		// 通过传递 zap.AddCallerSkip(n) 来调整 caller 的偏移量。
		// 这个选项可以告诉 zap 跳过一定数量的堆栈帧，以获取准确的调用者信息
		zapLogger = zapLogger.WithOptions(zap.AddCallerSkip(3))
	}
	return &ZapGormLogger{
		logger: zapLogger,
		config: lcf,
	}
}

func (l *ZapGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newConfig := l.config
	newConfig.LogLevel = level
	return NewZapGormLogger(l.logger, newConfig)
}

func (l *ZapGormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.config.LogLevel >= logger.Info {
		l.logger.Info(msg)
	}
}

func (l *ZapGormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.config.LogLevel >= logger.Warn {
		l.logger.Warn(msg)
	}
}

func (l *ZapGormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.config.LogLevel >= logger.Error {
		l.logger.Error(msg)
	}
}

func (l *ZapGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.config.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	msg := fmt.Sprintf("\r\n[%s] [rows:%d] %s", elapsed.String(), rows, sql)

	switch {
	case err != nil && l.config.LogLevel >= logger.Error:
		l.logger.Error(msg, zap.Error(err))
	case elapsed > l.config.SlowThreshold && l.config.SlowThreshold != 0 && l.config.LogLevel >= logger.Warn:
		l.logger.Warn(msg)
	case l.config.LogLevel >= logger.Info:
		l.logger.Info(msg)
	}
}
