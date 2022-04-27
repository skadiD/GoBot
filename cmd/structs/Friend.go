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
