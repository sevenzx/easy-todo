package main

import (
	"easytodo/config"
	"easytodo/global/vars"
	"easytodo/middleware"
	"easytodo/model"
	"easytodo/model/response"
	"easytodo/setup"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	setup.Viper()
	vars.Logger = setup.Zap()
	vars.DB = setup.GormMySQL()
	if vars.DB != nil {
		setup.RegisterTables()
		db, _ := vars.DB.DB()
		defer db.Close()
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestId())
	r.Use(middleware.AccessLog())
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		vars.Logger.Info("success", zap.String("msg", "pong"))
		var user model.User
		vars.DB.Where("id = ?", 1).Select("*").First(&user)
		response.Success(c, user)
	})

	_ = r.Run(config.Server.Port)
}
