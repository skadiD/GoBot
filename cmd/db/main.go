package db

import (
	"github.com/fexli/logger"
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
		logger.RootLogger.Error(logger.WithContent("数据库链接失败", err))
	}
	err = globals.Db.AutoMigrate(&ClassGroup{})
	err = globals.Db.AutoMigrate(&Session{})
	if err != nil {
		logger.RootLogger.Error(logger.WithContent("数据库链接失败", err))
	}
	sqlDB, _ := globals.Db.DB()
	sqlDB.SetConnMaxLifetime(7 * time.Hour)
	sqlDB.SetMaxIdleConns(200)
	logger.RootLogger.System(logger.WithContent("数据库链接成功"))
}
