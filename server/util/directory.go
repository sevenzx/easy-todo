package util

import (
	"github.com/pkg/errors"
	"os"
)

// DirExists 判断文件夹是否存在
func DirExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在和文件夹同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
