package structs

type BotJoinGroupEvent struct {
	SyncId string `json:"syncId"`
	Data   struct {
		Type    string   `json:"type"`
		Group   Group    `json:"group"`
		Invitor Operator `json:"invitor"`
	} `json:"data"`
}
type BotLeaveEventKick struct {
	SyncId string `json:"syncId"`
	Data   struct {
		Type     string   `json:"type"`
		Group    Group    `json:"group"`
		Operator Operator `json:"operator"`
	} `json:"data"`
}
type BotMuteEvent struct {
	SyncId string `json:"syncId"`
	Data   struct {
		Type            string   `json:"type"`
		DurationSeconds int      `json:"durationSeconds"`
		Operator        Operator `json:"operator"`
	} `json:"data"`
}

type BotUnmuteEvent struct {
	SyncId string `json:"syncId"`
	Data   struct {
		Type     string   `json:"type"`
		Operator Operator `json:"operator"`
	} `json:"data"`
}

type GroupMessage struct {
	SyncId string `json:"syncId"`
	Data   struct {
		Type         string         `json:"type"`
		MessageChain []MessageChain `json:"messageChain"`
		Sender       Operator       `json:"sender"`
	} `json:"data"`
}

type GroupRecallEvent struct {
	SyncId string `json:"syncId"`
	Data   struct {
		Type      string   `json:"type"`
		AuthorId  int      `json:"authorId"`
		MessageId int      `json:"messageId"`
		Time      int      `json:"time"`
		Group     Group    `json:"group"`
		Operator  Operator `json:"operator"`
	} `json:"data"`
}
type SendGroupMessage struct {
	SessionKey   string         `json:"sessionKey"`
	Target       int64          `json:"target"`
	MessageChain []MessageChain `json:"messageChain"`
}
