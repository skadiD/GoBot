package db

import "gobot/cmd/globals"

// GetGroup 判断是否是Master
func GetGroup(userId string) []ClassGroup {
	var cg []ClassGroup
	globals.Db.Where("master LIKE ?", "%"+userId+"%").Find(&cg)
	return cg
}

// FindGroup 查找group
func FindGroup(groupID int32) bool {
	var cg ClassGroup
	globals.Db.Where("group_id = ?", groupID).First(&cg)
	return cg.GroupID != 0
}

// CreateGroup 创建新班级
func CreateGroup(name, master, member string, groupID int32) {
	if !FindGroup(groupID) {
		cg := ClassGroup{
			Name:    name,
			Master:  master,
			Members: member,
			GroupID: groupID,
		}
		globals.Db.Create(&cg)
	}
}

// CreateSession 新增Session
func CreateSession(userID, groupID int32, target string, ts int) {
	s := Session{
		UserID:   userID,
		GroupID:  groupID,
		Command:  target,
		KeepTime: ts,
	}
	globals.Db.Create(&s)
}

// GetSession 获取session信息
func GetSession(userID, groupID int32, command string) *Session {
	s := Session{
		UserID:  userID,
		GroupID: groupID,
		Command: command,
	}
	globals.Db.First(&s)
	return &s
}

// DelSession 删除session信息
func DelSession(id any) {
	globals.Db.Delete(&Session{}, id)
}

// SetSession 设置 session 信息
func SetSession(userID, groupID int32, command, target, data string, keepTs int) {
	_old := GetSession(userID, groupID, command)
	_old.Target = target
	_old.Data = data
	_old.KeepTime = keepTs
	globals.Db.Model(&Session{}).Updates(_old)
}
