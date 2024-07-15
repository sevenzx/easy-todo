package router

import (
	"easytodo/core/api"
	"easytodo/middleware"
)

func (r *register) RegisterUserRouter() {
	basicRouter := r.BaseRouter.Group("/user")
	userRouter := r.BaseRouter.Group("/user")
	userRouter.Use(middleware.JWTAuth())
	{
		// 注册登录不需要JWT验证
		basicRouter.POST("/register", api.User.Register)
		basicRouter.POST("/login", api.User.Login)
	}
	{
		userRouter.GET("/info", api.User.UserInfo)
		userRouter.GET("/logout", api.User.Logout)
	}
}
