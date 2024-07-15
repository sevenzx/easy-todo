package api

import (
	"easytodo/core/service"
	"easytodo/global"
	"easytodo/model"
	"easytodo/model/req"
	"easytodo/model/resp"
	"easytodo/model/result"
	"easytodo/util/jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type userAPI struct{}

// Register 注册
func (api *userAPI) Register(c *gin.Context) {
	var params req.UserRegister
	_ = c.ShouldBindJSON(&params)
	err := service.User.Register(&params)
	if err != nil {
		result.Fail(c, err)
	} else {
		result.Ok(c)
	}
}

// Login 登录
func (api *userAPI) Login(c *gin.Context) {
	var params req.UserLogin
	_ = c.ShouldBindJSON(&params)
	user, err := service.User.Login(params.Username, params.Password)
	if err != nil {
		result.Fail(c, err)
		return
	}
	// 登录成功 签发jwt
	j := jwt.NewHelper()
	claims := j.CreateClaims(model.CustomClaims{
		UUID:     user.UUID,
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		result.Fail(c, err)
		return
	}
	// 向客户端设置token
	jwt.SetToken(c, token, int(claims.ExpiresAt.Unix()-time.Now().Unix()))
	result.OkWithData(c, resp.UserResp{
		User:      user,
		Token:     token,
		ExpiresAt: claims.ExpiresAt.Format(time.DateTime),
	})
}

// Logout 退出登录
func (api *userAPI) Logout(c *gin.Context) {
	service.User.Logout(c)
	result.Ok(c)
}

// UserInfo 通过c获取登录用户的信息
func (api *userAPI) UserInfo(c *gin.Context) {
	claims := jwt.GetClaims(c)
	u, err := service.User.GetUserByUuid(claims.UUID)
	if err != nil {
		global.Logger.Error("User Info", zap.Error(err))
		result.Fail(c, err)
	} else {
		result.OkWithData(c, u)
	}
}
