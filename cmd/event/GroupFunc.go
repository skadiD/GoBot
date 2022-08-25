package event

import (
	"github.com/fexli/logger"
	"github.com/goccy/go-json"
	"github.com/tidwall/gjson"
	"gobot/cmd/structs"
	"strings"
)

// BotJoinGroupEvent 邀请BOT入群
func BotJoinGroupEvent(_ *gjson.Result, origin []byte) {
	var _data structs.BotJoinGroupEvent
	json.Unmarshal(origin, &_data)
	///globals.Logger().Info("好友 " + strconv.FormatInt(_data.Data.Invitor.ID, 10) + " 邀请进群").
	///Success("群号: " + strconv.FormatInt(_data.Data.Group.ID, 10) + "\t群名: " + _data.Data.Group.Name).Run()
	logger.RootLogger.Warning(logger.WithContent("好友",
		_data.Data.Invitor.ID, "邀请进群",
	))
}

// BotLeaveEventKick BOT被踢群
func BotLeaveEventKick(_ *gjson.Result, origin []byte) {
	var _data structs.BotLeaveEventKick
	json.Unmarshal(origin, &_data)
	logger.RootLogger.Notice(logger.WithContent("群员",
		_data.Data.Operator.ID, "踢除BOT",
	))
}

// BotMuteEvent BOT被禁言
func BotMuteEvent(_ *gjson.Result, origin []byte) {
	var _data structs.BotMuteEvent
	json.Unmarshal(origin, &_data)
	logger.RootLogger.Warning(logger.WithContent("群员",
		_data.Data.Operator.ID, "禁言BOT",
	))
}

// BotUnmuteEvent BOT被解除禁言
func BotUnmuteEvent(_ *gjson.Result, origin []byte) {
	var _data structs.BotUnmuteEvent
	json.Unmarshal(origin, &_data)
	logger.RootLogger.Warning(logger.WithContent("群员",
		_data.Data.Operator.ID, "取消禁言BOT",
	))
	//globals.Logger().
	//	Warn("群员 " + strconv.FormatInt(_data.Data.Operator.ID, 10) + " 取消禁言BOT").
	//	Success("群号: " + strconv.FormatInt(_data.Data.Operator.Group.ID, 10) + "\t群名: " + _data.Data.Operator.Group.Name).Run()
}

// GroupMessage 群聊消息
// 解析捏麻麻
func GroupMessage(data *gjson.Result, origin []byte) (string, int64) {
	chain := data.Get("data.messageChain.1")
	group := data.Get("data.sender")
	switch chain.Get("type").String() {
	case "File":
		logger.RootLogger.Notice(logger.WithContent("群",
			group.Get("group.id").String(), "内成员",
			group.Get("id").String(), "上传文件",
			chain.Get("name").String()),
		)
	case "Plain":
		logger.RootLogger.Notice(logger.WithContent("群",
			group.Get("group.id").String(), "内成员",
			group.Get("id").String(), "发送消息",
			chain.Get("text").String()),
		)
		return strings.Trim(chain.Get("text").String(), " "), group.Get("group.id").Int()
	case "Image":
		logger.RootLogger.Notice(logger.WithContent("群",
			group.Get("group.id").String(), "内成员",
			group.Get("id").String(), "发送照片",
			chain.Get("url").String()),
		)
	}
	return "", 0
}
