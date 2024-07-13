// 全局变量

package vars

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Logger *zap.Logger
	DB     *gorm.DB
)
