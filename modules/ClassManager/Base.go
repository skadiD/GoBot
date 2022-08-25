package ClassManager

import (
	"github.com/tidwall/gjson"
	"gobot/cmd/consts"
	"gobot/cmd/event"
	"gobot/cmd/globals"
)

func Main() {
	event.AddHook(consts.GroupMessage, GroupMessage)
	//event.AddHook(consts.FriendMessage, FriendMessage)
}
func GroupMessage(data *gjson.Result, origin []byte) {
	str, _ := event.GroupMessage(data, origin)
	if str == "帮助" {
		globals.Bot.SendGroupMessage(677098563, "欢迎使用自助服务BOT\n当前实装功能\n1.绍兴自助过白\n\n")
	}
}
func FriendMessage(data *gjson.Result, origin []byte) {
	//str := event.GroupMessage(data, origin)
	//friend := data.Get("data.sender.id").Int()
}
