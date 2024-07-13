package internal

import (
	"fmt"
)

type MySQL struct {
	Path          string `mapstructure:"path" json:"path" yaml:"path"`
	Port          string `mapstructure:"port" json:"port" yaml:"port"`
	Config        string `mapstructure:"config" json:"config" yaml:"config"`                         // 高级配置
	Dbname        string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`                         // 数据库名
	Username      string `mapstructure:"username" json:"username" yaml:"username"`                   // 数据库用户名
	Password      string `mapstructure:"password" json:"password" yaml:"password"`                   // 数据库密码
	MaxIdleConns  int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns  int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 数据库的最大连接数
	Engine        string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"`        // 数据库引擎，默认InnoDB
	TablePrefix   string `mapstructure:"table-prefix" json:"table-prefix" yaml:"table-prefix"`       // 表前缀
	SingularTable bool   `mapstructure:"singular-table" json:"singular-table" yaml:"singular-table"` // 是否开启全局禁用复数
	LogLevel      string `mapstructure:"log-level" json:"log-level" yaml:"log-level"`                // 日志级别
	UseZap        bool   `mapstructure:"use-zap" json:"use-zap" yaml:"use-zap"`                      // 是否使用zap
}

// Dsn 获取dsn data source name
// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name
func (m *MySQL) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", m.Username, m.Password, m.Path, m.Port, m.Dbname, m.Config)
}
