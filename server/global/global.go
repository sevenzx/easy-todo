// 全局变量

package global

import (
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	Logger *zap.Logger
	DB     *gorm.DB
	Once   = &singleflight.Group{}
)
