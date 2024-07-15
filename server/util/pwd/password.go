package pwd

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// Generate 使用 bcrypt 对密码进行加密
func Generate(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// Check 对比明文密码和数据库的哈希值
func Check(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println("err", err)
	return err == nil
}
