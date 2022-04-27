package command

import "gobot/cmd/structs"

// SendFriendMessage 发送好友消息
func (b Bot) SendFriendMessage(friend int64, msg string) {
	test := structs.SendGroupMessage{ // 复用群消息发送结构体
		SessionKey: "",
		Target:     friend,
		MessageChain: []structs.MessageChain{{
			Type: "Plain",
			Text: msg,
		}},
	}
	Send("sendFriendMessage", test, b.Ws)
}
