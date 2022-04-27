package structs

type Base struct {
	Type string `json:"type"`
	QQ   int64  `json:"qq"`
}
type Friend struct {
	ID       int64  `json:"id"`       // 好友QQ号
	NickName string `json:"nickname"` // 好友昵称
	Remark   string `json:"remark"`   // 好友备注
}
type Group struct {
	ID         int64  `json:"id"`         // 群号
	Name       string `json:"name"`       // 群名
	Permission string `json:"permission"` // BOT在本群权限 OWNER或ADMINISTRATOR
}
type Operator struct {
	ID                 int64  `json:"id"`         // 操作员ID
	MemberName         string `json:"memberName"` // 操作员QQ号
	Permission         string `json:"permission"` // 操作员在本群权限
	Group              Group  `json:"group"`
	SpecialTitle       string `json:"specialTitle"`       // 群头衔
	JoinTimestamp      uint32 `json:"joinTimestamp"`      // 加入时间
	LastSpeakTimestamp uint32 `json:"lastSpeakTimestamp"` // 最后发言时间
	MuteTimeRemaining  uint32 `json:"muteTimeRemaining"`  // 禁言持续时间
}
type MessageChain struct {
	Type string `json:"type"`
	Id   any    `json:"id"`
	Time int    `json:"time"`
	Size int    `json:"size"`
	Name string `json:"name"`
	Text string `json:"text"`
}
type SendBase struct {
	SyncId     int         `json:"syncId"`
	Command    string      `json:"command"`
	SubCommand interface{} `json:"subCommand"`
	Content    interface{} `json:"content"`
}
