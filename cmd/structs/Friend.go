package structs

type FriendMessage struct {
	SyncId string `json:"syncId"`
	Data   struct {
		Type         string `json:"type"`
		MessageChain []struct {
			Type string `json:"type"`
			Id   int    `json:"id"`
			Time int    `json:"time"`
			Text string `json:"text"`
		} `json:"messageChain"`
		Sender struct {
			Id       int    `json:"id"`
			Nickname string `json:"nickname"`
			Remark   string `json:"remark"`
		} `json:"sender"`
	} `json:"data"`
}
type FriendRecallEvent struct {
	SyncId string `json:"syncId"`
	Data   struct {
		Type      string `json:"type"`
		AuthorId  int    `json:"authorId"`
		MessageId int    `json:"messageId"`
		Time      int    `json:"time"`
		Operator  int    `json:"operator"`
	} `json:"data"`
}

type FriendRequestEvent struct {
	SyncId string `json:"syncId"`
	Data   struct {
		Type    string `json:"type"`
		EventId int    `json:"eventId"`
		FromId  int    `json:"fromId"`
		GroupId int    `json:"groupId"`
		Nick    string `json:"nick"`
		Message string `json:"message"`
	} `json:"data"`
}
type ExecFriendReq struct {
	SessionKey string `json:"sessionKey"`
	EventId    int    `json:"eventId"`
	FromId     int    `json:"fromId"`
	GroupId    int    `json:"groupId"`
	Operate    int    `json:"operate"`
	Message    string `json:"message"`
}
