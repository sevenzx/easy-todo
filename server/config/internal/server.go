package internal

import "strings"

type Server struct {
	Port       string `mapstructure:"port" json:"port" yaml:"port"`                      // 服务监听端口
	BaseRouter string `mapstructure:"base-router" json:"base-router" yaml:"base-router"` // 服务基础路由
}

func (s *Server) CheckHostPort() {
	if !strings.HasPrefix(s.Port, ":") {
		s.Port = ":" + s.Port
	}
}
