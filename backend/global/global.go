package global

import (
	"gorm.io/gorm"
	"time"
)

// NOTE WARNING!!!
// Not the end code,maybe will change to iota
const MinUserPermissionCode = -1 // admin Permission Code
const MaxUserPermissionCode = 0  // normal Permission Code

// User Permission Code
const (
	AdminUser = iota - 1
	NormalUser
)

const ThreeDays = 3 * 24 * time.Hour

// JwtKey is the key used to sign JWT tokens
var JwtKey []byte

// DB is the global database
var DB *gorm.DB

// BcryptCost is the cost factor used by bcrypt
const BcryptCost = 12
