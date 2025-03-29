package permission

import (
	"FashOJ_Backend/models"
)

const (
	CreateProblem uint32 = 1 << iota
	ModifyProblem
	DelProblem

	CreateContest
	ModifyContest
	DelContest

	CreateAnnouncement
	ModifyAnnouncement
	DelAnnouncement

	ModifyPermission

	All uint32 = CreateProblem | ModifyProblem | DelProblem |
		CreateContest | ModifyContest | DelContest |
		CreateAnnouncement | ModifyAnnouncement | DelAnnouncement |
		ModifyPermission
)

func HasPermission(user models.User, permission uint32) bool {
	return (user.Permission | permission) != 0
}

func IsVaild(permission uint32) bool {
	// 有效权限掩码（包含所有预定义权限位）
	validMask := All

	// 检查权限位是否全部在预定义范围内
	return (permission & ^validMask) == 0
}
