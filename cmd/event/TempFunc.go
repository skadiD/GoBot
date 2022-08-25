package event

import (
	"github.com/fexli/logger"
	"github.com/tidwall/gjson"
	"strings"
)

func TempMessage(data *gjson.Result) (string, int64, int64) {
	chain := data.Get("data.messageChain.1")
	sender := data.Get("data.sender")
	user := sender.Get("id").Int()
	formGroup := sender.Get("group.id").Int()
	message := chain.Get("text").String()
	switch chain.Get("type").String() {
	case "Plain":
		logger.RootLogger.Notice(logger.WithContent("群 ", formGroup, " 内成员 ", user, " 发送私聊消息", message))
		return strings.Trim(message, " "), user, formGroup
	}
	return "", 0, 0
}
