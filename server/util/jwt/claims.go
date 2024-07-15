package jwt

import (
	"easytodo/global/consts"
	"easytodo/model"
	"github.com/gin-gonic/gin"
	"net"
)

// SetToken 设置Token
func SetToken(c *gin.Context, token string, maxAge int) {
	// 增加cookie x-token
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		// "/"：cookie的路径，表示cookie在整个域中都有效
		// ""：cookie的域名，留空表示只在当前域有效
		c.SetCookie(consts.JwtTokenKey, token, maxAge, "/", "", false, false)
	} else {
		c.SetCookie(consts.JwtTokenKey, token, maxAge, "/", host, false, false)
	}
}

// GetToken 获取Token
func GetToken(c *gin.Context) string {
	token, _ := c.Cookie(consts.JwtTokenKey)
	if token == "" {
		token = c.GetHeader(consts.JwtTokenKey)
	}
	return token
}

// ClearToken 清除Token
func ClearToken(c *gin.Context) {
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	if net.ParseIP(host) != nil {
		c.SetCookie(consts.JwtTokenKey, "", -1, "/", "", false, false)
	} else {
		c.SetCookie(consts.JwtTokenKey, "", -1, "/", host, false, false)
	}
}

// GetClaims 获取Claims
func GetClaims(c *gin.Context) *model.Claims {
	value, exists := c.Get(consts.JwtClaimsKey)
	if !exists {
		claims, err := GetClaimsFormToken(c)
		if err != nil {
			return nil
		} else {
			return claims
		}
	}
	claims, ok := value.(*model.Claims)
	if !ok {
		return nil
	} else {
		return claims
	}
}

// GetClaimsFormToken 从token中获取claims
func GetClaimsFormToken(c *gin.Context) (*model.Claims, error) {
	token := GetToken(c)
	j := NewHelper()
	claims, err := j.ParseToken(token)
	if err != nil {
		return nil, err
	} else {
		return claims, nil
	}
}
