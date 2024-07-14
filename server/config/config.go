package config

// yaml yaml配置文件
type yaml struct {
	Server     server     `mapstructure:"server" json:"server" yaml:"server"`
	MySQL      mysql      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	JWT        jwt        `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap        zap        `mapstructure:"zap" json:"zap" yaml:"zap"`
	Lumberjack lumberjack `mapstructure:"lumberjack" json:"lumberjack" yaml:"lumberjack"`
}

// Check 检查配置参数是否正确
func (y *yaml) Check() {
	y.Server.CheckHostPort()
}

var Yaml = new(yaml) // Viper加载时使用

// GlobalConfig 全局配置
var (
	Server     = &Yaml.Server
	MySQL      = &Yaml.MySQL
	JWT        = &Yaml.JWT
	Zap        = &Yaml.Zap
	Lumberjack = &Yaml.Lumberjack
)
