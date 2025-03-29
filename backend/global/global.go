package global

import (
	"os"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

const ValidTime = 3 * 24 * time.Hour

// JwtKey is the key used to sign JWT tokens
var JwtKey []byte

// DB is the global database
var DB *gorm.DB

// BcryptCost is the cost factor used by bcrypt
const BcryptCost = 12

// SystemTempFolder is the path to the system temporary folder
var SystemTempFolder = os.TempDir()

var Logger *zap.SugaredLogger