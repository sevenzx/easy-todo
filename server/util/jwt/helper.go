package jwt

import (
	"easytodo/config"
	"easytodo/global"
	"easytodo/model"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"strings"
	"time"
)

type Helper struct {
	SigningKey []byte
}

func NewHelper() *Helper {
	return &Helper{
		SigningKey: []byte(config.JWT.SigningKey),
	}
}

// CreateClaims 创建一个Claims
func (h *Helper) CreateClaims(customClaims model.CustomClaims) model.Claims {
	bf, _ := ParseDuration(config.JWT.BufferTime)
	ep, _ := ParseDuration(config.JWT.ExpiresTime)
	claims := model.Claims{
		//
		CustomClaims: customClaims,
		// 缓冲时间
		// 距离过期时间在缓冲时间内会获得新的token刷新令牌
		// 此时一个用户会存在两个有效令牌 但是前端只留一个(前端需要处理) 另一个会丢失
		BufferTime: int64(bf / time.Second),
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{customClaims.Username},   // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),    // 过期时间
			Issuer:    config.JWT.Issuer,                         // 签名的发行者
		},
	}
	return claims
}

// CreateToken 创建一个token
func (h *Helper) CreateToken(claims model.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(h.SigningKey)
}

// RefreshToken 使用之前的token来刷新token
func (h *Helper) RefreshToken(oldToken string, claims model.Claims) (string, error) {
	// 避免并发问题
	v, err, _ := global.Once.Do("JWT:"+oldToken, func() (interface{}, error) {
		return h.CreateToken(claims)
	})
	return v.(string), err
}

// ParseToken 解析token
func (h *Helper) ParseToken(t string) (*model.Claims, error) {
	token, err := jwt.ParseWithClaims(t, &model.Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return h.SigningKey, nil
	})
	if err != nil {
		return nil, err
	} else {
		if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
			return claims, nil
		}
		return nil, jwt.ErrTokenInvalidClaims
	}
}

func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour)
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}
