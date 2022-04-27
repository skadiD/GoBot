package db

import (
	"gobot/cmd/globals"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

func Conn() {
	var err interface{}
	globals.Db, err = gorm.Open(sqlite.Open(globals.Setting.DB), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		globals.Logger().Danger("数据库链接失败").Run()
	}
	err = globals.Db.AutoMigrate(&ClassGroup{})
	err = globals.Db.AutoMigrate(&Session{})
	if err != nil {
		globals.Logger().Danger("数据库表初始化失败").Run()
	}
	sqlDB, _ := globals.Db.DB()
	sqlDB.SetConnMaxLifetime(7 * time.Hour) // 比8小时最大值略小(线上环境)
	sqlDB.SetMaxIdleConns(200)              // 连接池最大链接
	globals.Logger().Success("数据库链接成功").Run()
}
