package event

import (
	"github.com/goccy/go-json"
	"github.com/tidwall/gjson"
	"gobot/cmd/globals"
	"gobot/cmd/structs"
	"strconv"
)

// https://segmentfault.com/a/1190000020336889
var eventMap = []func(data *gjson.Result, origin []byte){nil}

// TODO: 使用Int key
var eventIDMap = map[string]int{}

// AddHook 新增挂载点
func AddHook(target string, handle func(data *gjson.Result, origin []byte)) {
	eventIDMap[target] = len(eventMap)
	eventMap = append(eventMap, handle)
	globals.Logger().Warn(target + " 挂载成功").Run()
}

// BotOnlineEvent BOT登录成功
func BotOnlineEvent(data *gjson.Result, _ []byte) {
	globals.Logger().Success("Bot登录成功").Info("当前登录账号: " + data.Get("data.qq").String()).Run()
}

// BotOfflineEvent BOT离线
func BotOfflineEvent(data *gjson.Result, _ []byte) {
	globals.Logger().Warn("Bot离线").Info("退出原因: " + data.Get("data.type").String()).Run()
	globals.Logger().Warn("等待重连...").Run()
}

// BotReloginEvent BOT重登录
func BotReloginEvent(_ *gjson.Result, _ []byte) {
	globals.Logger().Warn("Bot正在重连, 请稍等").Run()
}

// FriendInputStatusChangedEvent 好友内容编辑状态
func FriendInputStatusChangedEvent(data *gjson.Result, _ []byte) {
	if data.Get("data.inputting").Bool() {
		globals.Logger().Info("好友 " + data.Get("data.friend.id").String() + " 正在编辑").Run()
	} else {
		globals.Logger().Info("好友 " + data.Get("data.friend.id").String() + " 结束编辑").Run()
	}
}

// FriendMessage 好友私信
func FriendMessage(_ *gjson.Result, origin []byte) (string, int) {
	var _data structs.FriendMessage
	json.Unmarshal(origin, &_data)
	globals.Logger().Info("好友 " + strconv.Itoa(_data.Data.Sender.Id) + " 发送消息").Warn("消息内容： " + _data.Data.MessageChain[1].Text).Run()
	//fmt.Printf(utils.Byte2Str(origin))
	return _data.Data.MessageChain[1].Text, _data.Data.Sender.Id
}

// FriendRecallEvent 好友撤回消息
func FriendRecallEvent(_ *gjson.Result, origin []byte) {
	var _data structs.FriendRecallEvent
	json.Unmarshal(origin, &_data)
	globals.Logger().Info("好友 " + strconv.Itoa(_data.Data.Operator) + " 撤销消息").Run()
	//fmt.Printf(utils.Byte2Str(origin))
}
