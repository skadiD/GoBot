package consts

const (
	BotGroupPermissionChangeEvent = iota + 10
	BotMuteEvent                  = "BotMuteEvent"
	BotUnmuteEvent                = "BotUnmuteEvent"
	BotJoinGroupEvent             = "BotJoinGroupEvent"
	BotLeaveEventActive           = "BotLeaveEventActive"
	BotLeaveEventKick             = "BotLeaveEventKick" // 被踢出群
	GroupMessage                  = "GroupMessage"      // 群聊消息
	GroupRecallEvent
	NudgeEvent
	GroupNameChangeEvent
	GroupEntranceAnnouncementChangeEvent
	GroupMuteAllEvent
	GroupAllowAnonymousChatEvent
	MemberJoinEvent
	MemberLeaveEventKick
	MemberLeaveEventQuit
	MemberCardChangeEvent
	MemberSpecialTitleChangeEvent
	MemberPermissionChangeEvent
	MemberMuteEvent
	MemberUnmuteEvent
	MemberJoinRequestEvent
	BotInvitedJoinGroupRequestEvent
)
