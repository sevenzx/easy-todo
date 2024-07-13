package internal

import (
	"go.uber.org/zap/zapcore"
	"time"
)

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                            // 级别
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 日志前缀
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                         // 输出格式 json | line
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`       // 编码级别
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` // 栈名
	AddCaller     bool   `mapstructure:"add-caller" json:"add-caller" yaml:"add-caller"`             // 添加调用者的文件名和行号
	LogInConsole  bool   `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` // 输出控制台
}

// LogLevel 根据字符串转化为 zapcore.Level
func (z *Zap) LogLevel() zapcore.Level {
	level, err := zapcore.ParseLevel(z.Level)
	if err != nil {
		level = zapcore.DebugLevel
	}
	return level
}

// Encoder 设置日志格式
func (z *Zap) Encoder() zapcore.Encoder {
	conf := zapcore.EncoderConfig{
		TimeKey:       "time",
		NameKey:       "name",
		LevelKey:      "level",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: z.StacktraceKey,
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(z.Prefix + " " + t.Format("2006-01-02 15:04:05.000"))
		},
		EncodeLevel:    z.LevelEncoder(),
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	if z.Format == "json" {
		return zapcore.NewJSONEncoder(conf)
	}
	return zapcore.NewConsoleEncoder(conf)

}

// LevelEncoder 根据 EncodeLevel 返回 zapcore.LevelEncoder
func (z *Zap) LevelEncoder() zapcore.LevelEncoder {
	switch {
	case z.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器
		return zapcore.LowercaseLevelEncoder
	case z.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case z.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		return zapcore.CapitalLevelEncoder
	case z.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色(默认)
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.CapitalColorLevelEncoder
	}
}
