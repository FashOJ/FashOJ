package global

import "gorm.io/gorm"

const MinUserRightCode = 0
const MaxUserRightCode = 1

const (
	NormalUser = iota
	AdminUser
)

var JwtKey []byte

var DB *gorm.DB