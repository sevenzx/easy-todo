package service

import (
	"easytodo/global"
	"easytodo/model"
	"easytodo/model/req"
	"easytodo/model/result/errcode"
	"easytodo/util/jwt"
	"easytodo/util/pwd"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type userService struct{}

// Register 注册
func (s *userService) Register(params *req.UserRegister) error {
	var user *model.User
	if !errors.Is(global.DB.Where("username = ?", params.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errcode.UserAlreadyExist
	}
	// 没有找到记录所以可以创建
	user.UUID, _ = uuid.NewRandom()
	user.Username = params.Username
	user.Nickname = params.Nickname
	user.Password = pwd.Generate(params.Password)
	err := global.DB.Create(&user).Error
	if err != nil {
		return errors.Wrap(err, "create user")
	} else {
		return nil
	}
}

// Login 用户登录
func (s *userService) Login(username string, password string) (*model.User, error) {
	var user model.User
	global.Logger.Info("user", zap.String("username", username), zap.String("password", password))
	err := global.DB.Where("username = ?", username).First(&user).Error
	if err == nil {
		if ok := pwd.Check(password, user.Password); !ok {
			return nil, errcode.PasswordError
		} else {
			return &user, nil
		}
	} else {
		return nil, errcode.UserNotExist
	}
}

// Logout 退出登录
func (s *userService) Logout(c *gin.Context) {
	jwt.ClearToken(c)
}

// GetUserByUuid 通过uuid获取用户信息
func (s *userService) GetUserByUuid(id uuid.UUID) (*model.User, error) {
	var user model.User
	err := global.DB.Where("uuid = ?", id).First(&user).Error
	return &user, err
}
