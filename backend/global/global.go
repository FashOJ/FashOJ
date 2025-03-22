package global

import "gorm.io/gorm"

const MinUserPermissionCode = -1
const MaxUserPermissionCode = 0

const (
	AdminUser = iota - 1
	NormalUser
)

var JwtKey []byte

var DB *gorm.DB