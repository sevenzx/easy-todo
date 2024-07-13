package setup

import (
	"easytodo/config"
	"easytodo/util"
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"time"
)

func Lumberjack() *lumberjack.Logger {
	// 可定制的输出目录。
	if ok, _ := util.DirExists(config.Lumberjack.Directory); !ok { // 判断是否有指定文件夹
		fmt.Printf("create %v directory\n", config.Lumberjack.Directory)
		_ = os.Mkdir(config.Lumberjack.Directory, os.ModePerm)
	}

	// 将文件名设置为日期
	logFileName := time.Now().Format(time.DateOnly) + ".log"
	fileName := path.Join(config.Lumberjack.Directory, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		// 不存在就创建
		if _, err := os.Create(fileName); err != nil {
			panic(err)
		}
	}
	// [配置详情](https://github.com/natefinch/lumberjack/tree/v2.0)
	return &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    config.Lumberjack.MaxSize,    // 一个文件最大可达?M。
		MaxBackups: config.Lumberjack.MaxBackups, // 最多同时保存?个文件。
		MaxAge:     config.Lumberjack.MaxAge,     // 一个文件最多可以保存?天。
		Compress:   config.Lumberjack.Compress,   // 是否用 gzip 压缩。
	}
}
