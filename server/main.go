package main

import (
	"easytodo/config"
	"easytodo/global"
	"easytodo/middleware"
	"easytodo/model"
	"easytodo/model/response"
	"easytodo/setup"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	setup.Viper()
	global.Logger = setup.Zap()
	global.DB = setup.GormMySQL()
	if global.DB != nil {
		setup.RegisterTables()
		db, _ := global.DB.DB()
		defer db.Close()
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestId())
	r.Use(middleware.AccessLog())
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		global.Logger.Info("success", zap.String("msg", "pong"))
		var user model.User
		global.DB.Where("id = ?", 1).Select("*").First(&user)
		response.Success(c, user)
	})

	_ = r.Run(config.Server.Port)
}
