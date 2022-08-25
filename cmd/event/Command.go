package event

import (
	"github.com/fexli/logger"
	"github.com/goccy/go-json"
	"github.com/tidwall/gjson"
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
	logger.RootLogger.System(logger.WithContent("新增挂载点: ", target))
}

// BotOnlineEvent BOT登录成功
func BotOnlineEvent(data *gjson.Result, _ []byte) {
	logger.RootLogger.System(logger.WithContent("Bot登录成功, 当前登录账号: ", data.Get("data.qq").String()))
}

// BotOfflineEvent BOT离线
func BotOfflineEvent(data *gjson.Result, _ []byte) {
	logger.RootLogger.Warning(logger.WithContent("Bot离线, 退出原因: ", data.Get("data.type").String()))
	logger.RootLogger.Warning(logger.WithContent("等待重连中..."))
}

// BotReloginEvent BOT重登录
func BotReloginEvent(_ *gjson.Result, _ []byte) {
	logger.RootLogger.Warning(logger.WithContent("BOT正在重连中..."))
}

// FriendInputStatusChangedEvent 好友内容编辑状态
func FriendInputStatusChangedEvent(data *gjson.Result, _ []byte) {
	if data.Get("data.inputting").Bool() {
		logger.RootLogger.Notice(logger.WithContent("好友 ", data.Get("data.friend.id").String(), " 正在输入..."))
	} else {
		logger.RootLogger.Notice(logger.WithContent("好友 ", data.Get("data.friend.id").String(), " 结束输入"))
	}
}

// FriendMessage 好友私信
func FriendMessage(_ *gjson.Result, origin []byte) (string, int) {
	var _data structs.FriendMessage
	json.Unmarshal(origin, &_data)
	logger.RootLogger.Notice(logger.WithContent("好友", _data.Data.Sender.Id, " 发送消息", _data.Data.MessageChain[1].Text))
	return _data.Data.MessageChain[1].Text, _data.Data.Sender.Id
}

// FriendRecallEvent 好友撤回消息
func FriendRecallEvent(_ *gjson.Result, origin []byte) {
	var _data structs.FriendRecallEvent
	json.Unmarshal(origin, &_data)
	logger.RootLogger.Notice(logger.WithContent("好友", _data.Data.Operator, " 撤销消息"))
}

// FriendRequestEvent 好友请求
func FriendRequestEvent(_ *gjson.Result, origin []byte) structs.FriendRequestEvent {
	var _data structs.FriendRequestEvent
	json.Unmarshal(origin, &_data)
	logger.RootLogger.Notice(logger.WithContent(
		"陌生人",
		_data.Data.Nick+"("+strconv.Itoa(_data.Data.FromId)+") 请求添加为好友",
		"添加理由: ", _data.Data.Message))
	return _data
}
