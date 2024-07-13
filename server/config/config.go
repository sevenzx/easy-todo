package config

type config struct {
	Server     server     `mapstructure:"server" json:"server" yaml:"server"`
	MySQL      mysql      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	JWT        jwt        `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap        zap        `mapstructure:"zap" json:"zap" yaml:"zap"`
	Lumberjack lumberjack `mapstructure:"lumberjack" json:"lumberjack" yaml:"lumberjack"`
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
