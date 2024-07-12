package internal

import (
	"easytodo/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type ZapCore struct {
	level zapcore.Level
	zapcore.Core
}

func NewZapCore(level zapcore.Level) *ZapCore {
	zapCore := &ZapCore{level: level}
	syncer := zapCore.WriteSyncer()
	levelEnabler := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == level
	})
	zapCore.Core = zapcore.NewCore(config.Zap.Encoder(), syncer, levelEnabler)
	return zapCore
}

func (z *ZapCore) WriteSyncer(formats ...string) zapcore.WriteSyncer {
	cutter := NewCutter(
		config.Zap.Directory,
		z.level.String(),
		config.Zap.RetentionDay,
		CutterWithLayout(time.DateOnly),
		CutterWithFormats(formats...),
	)
	if config.Zap.LogInConsole {
		multiSyncer := zapcore.NewMultiWriteSyncer(os.Stdout, cutter)
		return zapcore.AddSync(multiSyncer)
	}
	return zapcore.AddSync(cutter)
}

func (z *ZapCore) Enabled(level zapcore.Level) bool {
	return z.level == level
}

func (z *ZapCore) With(fields []zapcore.Field) zapcore.Core {
	return z.Core.With(fields)
}

func (z *ZapCore) Check(entry zapcore.Entry, check *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if z.Enabled(entry.Level) {
		return check.AddCore(entry, z)
	}
	return check
}

func (z *ZapCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	for i := 0; i < len(fields); i++ {
		// 命中key之后, 会根据value(fields[i].String)新建目录写入日志
		if fields[i].Key == "business" {
			syncer := z.WriteSyncer(fields[i].String)
			z.Core = zapcore.NewCore(config.Zap.Encoder(), syncer, z.level)
		}
	}
	return z.Core.Write(entry, fields)
}

func (z *ZapCore) Sync() error {
	return z.Core.Sync()
}
