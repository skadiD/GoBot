package core

import (
	"github.com/fexli/logger"
	"github.com/gorilla/websocket"
	"gobot/cmd/command"
	"gobot/cmd/event"
	"gobot/cmd/globals"
	"gobot/utils"
	"net/url"
	"os"
	"os/signal"
	"time"
)

var retry = 0

func Conn() {
	if retry >= 3 {
		logger.RootLogger.Error(logger.WithContent("尝试链接次数过多，已暂停链接"))
		return
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	u := url.URL{
		Scheme:   "ws",
		Host:     globals.Setting.Bot.Ws,
		Path:     "/all",
		RawQuery: "verifyKey=" + globals.Setting.Bot.VerifyKey,
	}
	logger.RootLogger.Notice(logger.WithContent("BOT正在链接中"))
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		logger.RootLogger.Error(logger.WithContent("与Mirai通讯失败", err))
		return
	}
	defer c.Close()
	done := make(chan struct{})
	logger.RootLogger.System(logger.WithContent("BOT链接成功"))
	globals.Bot = command.Bot{
		Ws: c,
	}
	go func() {
		defer close(done)
		for {
			_, message, readErr := c.ReadMessage()
			if readErr != nil {
				logger.RootLogger.Warning(logger.WithContent("BOT链接断开", readErr))
				retry++
				logger.RootLogger.Notice(logger.WithContent("正在尝试重连...第", retry, "次"))
				Conn()
				return
			}
			go event.ParseCore(message)
			if globals.Setting.Bot.Debug {
				logger.RootLogger.Debug(logger.WithContent("收到消息", utils.Byte2Str(message)))
			}
		}
	}()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			return
		case <-interrupt:
			logger.RootLogger.System(logger.WithContent("BOT正在退出..."))
			_err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if _err != nil {
				logger.RootLogger.Warning(logger.WithContent("BOT退出失败", _err))
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
