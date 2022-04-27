package core

import (
	"fmt"
	"github.com/gorilla/websocket"
	"gobot/cmd/command"
	"gobot/cmd/db"
	"gobot/cmd/event"
	"gobot/cmd/globals"
	"gobot/modules/ClassManager"
	"gobot/utils"
	"gobot/utils/Logger"
	"gobot/utils/config"
	"net/url"
	"os"
	"os/signal"
	"time"
)

func conn() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	u := url.URL{
		Scheme:   "ws",
		Host:     globals.Setting.Bot.Ws,
		Path:     "/all",
		RawQuery: "verifyKey=" + globals.Setting.Bot.VerifyKey,
	}
	globals.Logger().Info("BOT正在链接中").Run()

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		globals.Logger().Danger("与Mirai通讯失败").Info(err.Error()).Run()
		return
	}
	defer c.Close()
	done := make(chan struct{})
	globals.Logger().Success("BOT链接成功").Run()
	globals.Bot = command.Bot{
		Ws: c,
	}
	go func() {
		defer close(done)
		for {
			_, message, readErr := c.ReadMessage()
			if readErr != nil {
				globals.Logger().Warn("读取ws数据失败").Danger(readErr.Error()).Run()
				return
			}
			startT := time.Now()
			event.ParseCore(message)
			tc := time.Since(startT)
			fmt.Printf("cost: %v\n", tc)
			if globals.Setting.Bot.Debug {
				globals.Logger().Info("Debug").Success(utils.Byte2Str(message)).Run()
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
			globals.Logger().Info("BOT正在退出...").Run()
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				globals.Logger().Warn("BOT退出失败").Danger(err.Error()).Run()
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
func Fire() {
	globals.Logger = Logger.Create()
	globals.Setting = config.Read()
	db.Conn()
	ClassManager.Main()
	//.//globals.Logger().Warn("test").Info("test").Danger("test").Run()
	conn()
}
