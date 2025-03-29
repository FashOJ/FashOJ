package utils

import (
	"FashOJ_Backend/global"
	"FashOJ_Backend/models"
	"crypto/rand"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// GenJwt generates a JWT token for the given username.
func GenJwt(username string) (string, error) {
	// Create a new "token object", specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":       username,
		"expirationTime": time.Now().Add(global.ValidTime).Unix(),
	})
	// Use the token object's SignedString method to get the complete signed token string
	signedToken, err := token.SignedString(global.JwtKey)
	return "Bearer " + signedToken, err
}

// ParseJwt parses a JWT token string and returns the username if the token is valid.
func ParseJwt(tokenString string) (string, error) {

	// Remove the "Bearer " prefix from the token string
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// Parse the token string and extract the claims.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// Check the signing method and return the key used to sign the token.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexcept signing method")
		}
		return global.JwtKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { // Check if the token is valid and the claims are valid.
		username, ok := claims["username"].(string)
		if !ok {
			return "", errors.New("username claim is not a string")
		}
		return username, nil
	}
	return "", err
}

// Hash password with bcrypt.
func HashPwd(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), global.BcryptCost)
	return string(hash), err
}

// Check password with bcrypt.
func CheckPwd(pwd string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}

/*
HAC (Hash based Message Authentication Code) is a message authentication code based on hash functions,
used to verify the integrity and authenticity of messages.
*/
// GenerateHMACKey generates a random 256-bit key for HMAC signing.
func generateHMACKey() ([]byte, error) {
	key := make([]byte, 32)  // 256-bit key.
	_, err := rand.Read(key) // Fill key with random bytes.
	if err != nil {
		return []byte{}, err
	}
	return key, nil
}

func SetJwtKey() {
	keyString := os.Getenv("JWT_SECRET")
	if keyString == "" {
		key, err := generateHMACKey()
		if err != nil {
			panic(err)
		}
		global.JwtKey = key
	} else {
		global.JwtKey = []byte(keyString)
	}
}

//Auto Migrate all models
func AutoMigrate() {
	if err := global.DB.AutoMigrate(
		&models.User{},
		&models.Limit{},
		&models.Problem{},
		&models.Testcase{},
		&models.Announcement{},
	); err != nil {
		global.Logger.Panic(err.Error())
		panic(err)
	}
}
