package setup

import (
	"easytodo/config"
	"easytodo/setup/internal"
	"easytodo/util"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func Zap() (logger *zap.Logger) {
	if ok, _ := util.DirExists(config.Zap.Directory); !ok { // 判断是否有指定文件夹
		fmt.Printf("create %v directory\n", config.Zap.Directory)
		_ = os.Mkdir(config.Zap.Directory, os.ModePerm)
	}
	levels := config.Zap.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := internal.NewZapCore(levels[i])
		cores = append(cores, core)
	}
	logger = zap.New(zapcore.NewTee(cores...))
	if config.Zap.AddCaller {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
