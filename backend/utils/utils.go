package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GenJwt(username string) (string,error){
	secret := []byte("zineyu")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"username":username,
		"exp":time.Now().Add(time.Hour*72).Unix(),
	})

	signedToken,err := token.SignedString(secret)
	return "Bearer "+signedToken,err
}

func ParseJwt(tokenstring string) (string,error){
	if len(tokenstring) > 7 && tokenstring[:7] == "Bearer "{
		tokenstring = tokenstring[7:]
	}

	token,err := jwt.Parse(tokenstring,func (token *jwt.Token) (interface {},error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil,errors.New("unexcept signing method")
		}
		return []byte("secret"),nil
	})

	if err !=nil{
		return "",err
	}

	if claims,ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
		username,ok:=claims["username"].(string)

		if !ok {
			return "",errors.New("Username claim is not a string")

		}

		return username,nil
	}
	return "",err
}

func HashPwd(pwd string) (string,error){
	hash,err := bcrypt.GenerateFromPassword([]byte(pwd),12)
	return string(hash),err
}


func CheckPwd(pwd string,hash string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(pwd))
	return err == nil
}