package global

import "gorm.io/gorm"

const (
	NormalUser = 0
	AdminUser = 1
)

var DB *gorm.DB