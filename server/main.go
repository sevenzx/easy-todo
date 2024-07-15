package main

import (
	"easytodo/global"
	"easytodo/setup"
)

func main() {
	setup.Viper()
	global.Logger = setup.Zap()
	global.DB = setup.GormMySQL()
	if global.DB != nil {
		setup.RegisterTables()
		db, _ := global.DB.DB()
		defer db.Close()
	}
	setup.Gin()
}
