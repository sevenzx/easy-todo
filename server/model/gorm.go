package model

import (
	"gorm.io/gorm"
	"time"
)

type GORM struct {
	Id        uint           `gorm:"primarykey" json:"id"` // 主键Id
	CreatedAt time.Time      `json:"created_at"`           // 创建时间
	UpdatedAt time.Time      `json:"updated_at"`           // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`       // 删除时间
}
