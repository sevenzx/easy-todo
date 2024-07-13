package config

import "easytodo/config/internal"

type config struct {
	Server     internal.Server     `mapstructure:"server" json:"server" yaml:"server"`
	MySQL      internal.MySQL      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	JWT        internal.JWT        `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap        internal.Zap        `mapstructure:"zap" json:"zap" yaml:"zap"`
	Lumberjack internal.Lumberjack `mapstructure:"lumberjack" json:"lumberjack" yaml:"lumberjack"`
}

// Check 检查配置参数是否正确
func (c *config) Check() {
	c.Server.CheckHostPort()
}

var File = new(config) // Viper加载时使用

// GlobalConfig 全局配置
var (
	Server     = &File.Server
	MySQL      = &File.MySQL
	JWT        = &File.JWT
	Zap        = &File.Zap
	Lumberjack = &File.Lumberjack
)
