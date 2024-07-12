package middleware

import (
	"easytodo/global/consts"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
	"time"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		if requestId := c.Request.Header.Get(consts.RequestIdKey); requestId != "" {
			s := uuid.New().String()
			s = strings.ReplaceAll(s, "-", "")
			s = strings.ToUpper(s)
			requestId = fmt.Sprintf("%s%s", time.Now().Format("20060102150405"), s)
			c.Request.Header.Set(consts.RequestIdKey, requestId)
		}
		c.Next()
	}
}
