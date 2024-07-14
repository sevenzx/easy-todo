package middleware

import (
	"easytodo/global"
	"easytodo/global/consts"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func AccessLog() gin.HandlerFunc {
	notlogged := []string{""}
	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		// Process request
		c.Next()

		// Log only when it is not being skipped
		if _, ok := skip[path]; ok {
			return
		}
		if raw != "" {
			path = path + "?" + raw
		}
		global.Logger.Info("AccessLog",
			zap.String("request-id", c.Request.Header.Get(consts.RequestIdKey)),
			zap.String("latency", time.Now().Sub(start).String()),
			zap.String("client-ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.Int("status-code", c.Writer.Status()),
			zap.Int("body-size", c.Writer.Size()),
			zap.String("path", path),
		)
	}
}
