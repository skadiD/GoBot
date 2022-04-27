package command

import (
	"github.com/gorilla/websocket"
	"gobot/cmd/structs"
)

type Bot struct {
	Ws *websocket.Conn
}

// SendGroupMessage 发送群消息
func (b Bot) SendGroupMessage(group int64, msg string) {
	test := structs.SendGroupMessage{
		SessionKey: "",
		Target:     group,
		MessageChain: []structs.MessageChain{{
			Type: "Plain",
			Text: msg,
		}},
	}
	Send("sendGroupMessage", test, b.Ws)
}
