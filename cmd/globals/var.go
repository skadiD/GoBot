package globals

import (
	"gobot/cmd/command"
	"gorm.io/gorm"
)

var (
	Setting *Common
	Bot     command.Bot
	Db      *gorm.DB
)
