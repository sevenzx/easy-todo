package model

import "github.com/google/uuid"

type User struct {
	GORM
	UUID     uuid.UUID `json:"uuid" gorm:"index:,unique;comment:用户UUID"`       // 用户UUID
	Username string    `json:"username" gorm:"index:,unique;comment:用户名"`      // 用户登录名
	Password string    `json:"-"  gorm:"comment:密码"`                           // 用户登录密码
	Nickname string    `json:"nickname" gorm:"default:todo_user;comment:用户昵称"` // 用户昵称
}
