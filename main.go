package main

import (
	"github.com/fexli/logger"
	"gobot/cmd/core"
	"gobot/cmd/db"
	"gobot/cmd/globals"
	_ "gobot/modules/White"
	"gobot/utils/config"
)

func main() {
	logger.InitGlobLog("bot.log", "SkadiD Bot")
	globals.Setting = config.Read()
	db.Conn()
	core.Conn()
}
