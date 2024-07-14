package setup

import (
	"easytodo/config"
	"easytodo/global"
	"easytodo/model"
	"easytodo/setup/internal"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"strings"
	"time"
)

// GormMySQL 初始化Gorm的MySQL
func GormMySQL() *gorm.DB {
	if config.MySQL.Dbname == "" {
		return nil
	}
	gcf := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.MySQL.TablePrefix,
			SingularTable: config.MySQL.SingularTable,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	// 1. 设置日志等级
	var logLevel logger.LogLevel
	switch strings.ToLower(config.MySQL.LogLevel) {
	case "silent":
		logLevel = logger.Silent
	case "error":
		logLevel = logger.Error
	case "warn":
		logLevel = logger.Warn
	case "info":
		logLevel = logger.Info
	default:
		logLevel = logger.Info
	}

	// 2. 设置logger的配置
	lcf := logger.Config{
		SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
		LogLevel:                  logLevel,               // Log level
		IgnoreRecordNotFoundError: false,                  // Ignore ErrRecordNotFound error for logger
		ParameterizedQueries:      false,                  // Don't include params in the SQL log
		Colorful:                  false,                  // Disable color
	}
	if config.MySQL.UseZap {
		gcf.Logger = internal.NewZapGormLogger(
			global.Logger, // io writer
			lcf,
		)
	} else {
		lcf.Colorful = true
		gcf.Logger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			lcf,
		)
	}

	// [参考](https://gorm.io/zh_CN/docs/connecting_to_the_database.html)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.MySQL.Dsn(), // DSN data source name
		DefaultStringSize:         256,                // string 类型字段的默认长度
		SkipInitializeWithVersion: false,              // 根据当前 MySQL 版本自动配置
	}), gcf)
	if err != nil {
		global.Logger.Error("gorm.Open", zap.Error(err))
		return nil
	} else {
		// 设置数据库实例选项，指定表的存储引擎
		db.InstanceSet("gorm:table_options", "ENGINE="+config.MySQL.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(config.MySQL.MaxIdleConns)
		sqlDB.SetMaxOpenConns(config.MySQL.MaxOpenConns)
		return db
	}
}

// RegisterTables 初始化表
func RegisterTables() {
	db := global.DB
	err := db.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		global.Logger.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.Logger.Info("register table success")
}
