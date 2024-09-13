package setup

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"

	"easytodo/config"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Viper 使用viper加载配置文件
// 优先级: 命令行 > 默认值 > 自定义路径
func Viper(paths ...string) {
	var path string

	if len(paths) == 0 {
		flag.StringVar(&path, "c", "", "set config file path.")
		flag.Parse()
		if path != "" {
			// 1. 使用命令行
			fmt.Printf("Using cmd line file: %s\n", path)
		} else {
			// 2. 使用默认值
			path = "config.yaml"
			fmt.Printf("Using default file: %s\n", path)
		}
	} else {
		// 3. 使用自定义路径
		path = paths[0]
		fmt.Printf("Using custom file: %s\n", path)
	}

	v := viper.New()
	// 指定配置文件路径
	v.SetConfigFile(path)
	// 指定配置文件类型
	v.SetConfigType("yaml")
	// 读取配置信息
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error path file: %s \n", err))
	}
	// 监控配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("viper: config file(%s) changed", e.Name)
		if err = v.Unmarshal(&config.Yaml); err != nil {
			panic(errors.Wrap(err, "unmarshal config file error"))
		} else {
			// 检查配置参数是否填写正确
			config.Yaml.Check()
		}
	})
	// 将配置文件加载到config.Yaml中
	if err = v.Unmarshal(&config.Yaml); err != nil {
		panic(errors.Wrap(err, "unmarshal config file error"))
	} else {
		// 检查配置参数是否填写正确
		config.Yaml.Check()
	}
}
