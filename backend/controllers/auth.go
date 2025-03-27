package controllers

import (
	"FashOJ_Backend/global"
	"FashOJ_Backend/models"
	"FashOJ_Backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
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
			"Error": err,
		})
		return
	}

	// Find the user in the database
	var foundUser models.User
	if err := global.DB.Where("username = ?", UserRequestInput.Username).First(&foundUser).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Error": "wrong password or username",
		})
		return
	}

	// Check the password
	if !utils.CheckPwd(UserRequestInput.Password, foundUser.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Error": "wrong password or username",
		})
		return
	}

	// Generate the token
	token, err := utils.GenJwt(foundUser.Username)
	if err != nil {
		global.Logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": "server error",
		})
	}

	// Return the token
	ctx.JSON(http.StatusOK, gin.H{
		"Token": token,
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
			"Error": err,
		})
		return
	}

	// Hash the password
	hashedPwd, err := utils.HashPwd(NewUserRequest.Password)
	if err != nil {
		global.Logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": "server error",
		})
	}

	// Generate the token and return it to the user
	NewUserRequest.Password = hashedPwd
	token, err := utils.GenJwt(NewUserRequest.Username)
	if err != nil {
		global.Logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": "server error",
		})
		return
	}

	// Call the AutoMigrate method of global.DB to automatically migrate the database table structure.
	// Ensure that the table corresponding to the NewUserRequest struct exists; if not, create it.

	// Create the user into the database
	if err := global.DB.Create(&NewUserRequest).Error; err != nil {
		global.Logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": "server error",
		})
		return
	}

	// Return the token
	ctx.JSON(http.StatusOK, gin.H{
		"Token": token,
	})
}
