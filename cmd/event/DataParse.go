package event

import (
	"fmt"
	"github.com/tidwall/gjson"
	"gobot/cmd/globals"
)

// ParseCore 核心消息解析
func ParseCore(data []byte) {
	d := gjson.ParseBytes(data)
	eventTag := d.Get("data.type").String()
	if globals.Setting.Bot.Debug {
		fmt.Println(eventTag)
	}
	if eventTag != "" {
		if eventIDMap[eventTag] != 0 {
			eventMap[eventIDMap[eventTag]](&d, data)
			//} else {
			//globals.Logger().Warn("意外的PackTag -> " + eventTag).Danger(utils.Byte2Str(data)).Run()
		}
	}
}
