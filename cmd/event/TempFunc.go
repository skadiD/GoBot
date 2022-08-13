package event

import (
	"github.com/tidwall/gjson"
	"gobot/cmd/globals"
	"strconv"
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
		globals.Logger().
			Info("群 " + strconv.FormatInt(formGroup, 10) + " 内成员 " + strconv.FormatInt(user, 10) + " 发送私聊消息").
			Success("消息内容: " + message).
			Run()
		return strings.Trim(message, " "), user, formGroup
	}
	return "", 0, 0
}
