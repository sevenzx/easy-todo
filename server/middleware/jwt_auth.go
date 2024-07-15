package middleware

import (
	"easytodo/config"
	"easytodo/global/consts"
	"easytodo/model/response"
	"easytodo/model/response/errcode"
	jwtutil "easytodo/util/jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"time"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时会返回token信息
		// 这里前端需要把token存储到cookie或者本地localStorage中
		token := jwtutil.GetToken(c)
		// 1. 判断是否有token
		if token == "" {
			response.Fail(c, errcode.TokenIsNotExist)
			c.Abort()
			return
		}
		// 2. 验证token
		j := jwtutil.NewHelper()
		claims, err := j.ParseToken(token)
		if err != nil {
			// 如果token过期就清除客户端的token
			if errors.Is(err, jwt.ErrTokenExpired) {
				jwtutil.ClearToken(c)
			}
			response.Fail(c, errcode.TokenAuthFail)
			c.Abort()
			return
		}
		// 3. 在上下文中设置claims供后续使用
		c.Set(consts.JwtClaimsKey, claims)
		// 4. 判断是否需要刷新
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			duration, _ := jwtutil.ParseDuration(config.JWT.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(duration))
			newToken, _ := j.RefreshToken(token, *claims)
			c.Header("new-token", newToken)
			jwtutil.SetToken(c, newToken, int(duration.Seconds()))
		}
		c.Next()
	}
}
