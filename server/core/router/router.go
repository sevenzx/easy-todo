package router

import "github.com/gin-gonic/gin"

type register struct {
	BaseRouter *gin.RouterGroup
}

// RegisterV1Router 注册V1接口的路由
func RegisterV1Router(baseRouter *gin.RouterGroup) {
	v1Router := baseRouter.Group("/v1")
	r := register{
		BaseRouter: v1Router,
	}

	r.RegisterUserRouter()
}
