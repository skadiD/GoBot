package db

import "time"

type ClassGroup struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string // 群名
	GroupID   int32  `gorm:"index:group"` // 群号
	Master    string // 群管理员
	Members   string // 群成员
	CreatedAt time.Time
}
type Session struct {
	ID       uint   `gorm:"primaryKey"`
	UserID   int32  // 用户ID
	GroupID  int32  // 群ID 可空
	Command  string // 指令名称
	Target   string // 进度关键词
	Data     string // 存储数据
	KeepTime int    // 持续时间
}
