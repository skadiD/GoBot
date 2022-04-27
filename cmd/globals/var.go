package globals

import (
	"gobot/cmd/command"
	Logger2 "gobot/utils/Logger"
	"gorm.io/gorm"
)

var (
	Setting *Common
	Logger  func() *Logger2.Logger
	Bot     command.Bot
	Db      *gorm.DB
)
