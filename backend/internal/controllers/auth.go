package controllers

import (
	"FashOJ_Backend/internal/global"
	"FashOJ_Backend/internal/models"
	"FashOJ_Backend/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login is the controller for the login route.
// It takes the username and password from the request body and checks if the user exists in the database.
// If the user exists, it checks the password and returns the token.
// If the user does not exist, it returns an error.
func Login(ctx *gin.Context) {
	var UserRequestInput models.User

	// Bind the request body to the UserRequestInput struct
	if err := ctx.ShouldBindJSON(&UserRequestInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	// Find the user in the database
	var foundUser models.User
	if err := global.DB.Where("username = ?", UserRequestInput.Username).First(&foundUser).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "wrong password or username",
		})
		return
	}

	// Check the password
	if !utils.CheckPwd(UserRequestInput.Password, foundUser.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "wrong password or username",
		})
		return
	}

	// Generate the token
	token, err := utils.GenJwt(foundUser.Username)
	if err != nil {
		global.Logger.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
	}

	// Return the token

	global.Logger.Sugar().Infof("user %s login", foundUser.Username)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"token":   token,
	})
}

// Register is the controller for the register route.
// It takes the username and password from the request body and creates a new user in the database.
// Then returns the token.
func Register(ctx *gin.Context) {

	// Bind the request body to the NewUserRequest struct
	var NewUserRequest models.User
	if err := ctx.ShouldBindJSON(&NewUserRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	var findUser models.User
	if err := global.DB.
		Where("username = ?", NewUserRequest.Username).
		First(&findUser).Error; err != nil {
		global.Logger.Error(err.Error())
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "用户名重复",
		})
		return
	}

	// Hash the password
	hashedPwd, err := utils.HashPwd(NewUserRequest.Password)
	if err != nil {
		global.Logger.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}

	// Generate the token and return it to the user
	NewUserRequest.Password = hashedPwd
	token, err := utils.GenJwt(NewUserRequest.Username)
	if err != nil {
		global.Logger.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}

	// Call the AutoMigrate method of global.DB to automatically migrate the database table structure.
	// Ensure that the table corresponding to the NewUserRequest struct exists; if not, create it.

	// Create the user into the database
	if err := global.DB.Create(&NewUserRequest).Error; err != nil {
		global.Logger.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}

	// Return the token
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"token":   token,
	})
}
