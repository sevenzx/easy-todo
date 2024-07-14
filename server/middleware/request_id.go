package middleware

import (
	"easytodo/global/consts"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		if requestId := c.Request.Header.Get(consts.RequestIdKey); requestId == "" {
			requestId = uuid.New().String()
			requestId = strings.ReplaceAll(requestId, "-", "")
			requestId = strings.ToUpper(requestId)
			c.Request.Header.Set(consts.RequestIdKey, requestId)
		}
		c.Next()
	}
}
