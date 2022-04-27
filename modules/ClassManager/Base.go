package ClassManager

import (
	"github.com/tidwall/gjson"
	"gobot/cmd/consts"
	"gobot/cmd/db"
	"gobot/cmd/event"
	"gobot/cmd/globals"
)

func Main() {
	event.AddHook(consts.GroupMessage, GroupMessage)
	event.AddHook(consts.FriendMessage, FriendMessage)

}
func GroupMessage(data *gjson.Result, origin []byte) {
	str := event.GroupMessage(data, origin)
	if str == "帮助" {
		globals.Bot.SendGroupMessage(677098563, "欢迎使用智慧班务BOT\n当前开启的功能\n1.创建接龙任务\n2.创建收集任务\n\n回复数字以执行")
	}
}
func FriendMessage(data *gjson.Result, origin []byte) {
	str := event.GroupMessage(data, origin)
	friend := data.Get("data.sender.id").Int()
	if str == "帮助" {
		globals.Bot.SendFriendMessage(friend, "欢迎使用智慧班务BOT\n当前开启的功能\n1.创建接龙任务\n2.创建收集任务\n\n回复数字以执行")
	} else if str == "班级新增" {
		db.CreateSession(int32(friend), 0, "Create", 120)
		globals.Bot.SendFriendMessage(friend, "请输入班级名称")

		globals.Bot.SendGroupMessage(friend, "请输入班级群号")
		globals.Bot.SendGroupMessage(friend, "请输入班级管理员")
		db.CreateGroup("test", "1 2 3", "1 2 3", 10000)
		go globals.Bot.SendFriendMessage(friend, "创建成功")
		globals.Logger().Success("创建成功").Run()
	}
}
