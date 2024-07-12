package main

import (
	"easytodo/config"
	"easytodo/global/vars"
	"easytodo/middleware"
	"easytodo/setup"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	setup.Viper()
	vars.Logger = setup.Zap()

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestId())
	r.Use(middleware.AccessLog())
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		vars.Logger.Info("success", zap.String("msg", "pong"))
		c.String(http.StatusOK, "pong")
	})

	_ = r.Run(config.Server.Port)
}
