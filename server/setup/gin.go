package setup

import (
	"easytodo/config"
	"easytodo/core/router"
	"easytodo/middleware"
	"github.com/gin-gonic/gin"
)

func Gin() {
	// 1. 新建engine
	engine := gin.New()
	// 2. 注册中间件
	engine.Use(gin.Recovery())
	engine.Use(middleware.RequestId())
	engine.Use(middleware.AccessLog())

	// 3. 注册base路由
	baseRouter := engine.Group(config.Server.BaseRouter)
	router.RegisterV1Router(baseRouter)
	// 4. 启动服务
	_ = engine.Run(config.Server.Port)
}
