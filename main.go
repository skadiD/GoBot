package main

import (
	"github.com/fexli/logger"
	"gobot/cmd/core"
)

func main() {
	logger.InitGlobLog("bot.log", "SkadiD Bot")
	core.Fire()
}
