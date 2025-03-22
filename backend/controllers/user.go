package controllers

import (
	"FashOJ_Backend/global"
	"FashOJ_Backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChangeRight(ctx *gin.Context) {
	user := ctx.Value("user").(models.User)

	var input struct {
		Username string
		Right int
	}

	if err := ctx.ShouldBindJSON(&input); err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"error":err})
		return
	}

	if user.Right == global.AdminUser {
		var userChangeRight models.User
		if err:=global.DB.Where("username = ?",input.Username).First(&userChangeRight).Error; err != nil{
			ctx.JSON(http.StatusBadRequest,gin.H{"error":"no such user"})
			return
		}
		if input.Right < global.MinUserRightCode || input.Right > global.MaxUserRightCode{
			ctx.JSON(http.StatusBadRequest,gin.H{"error":"wrong right code"})
			return
		}
		userChangeRight.Right = input.Right
		global.DB.Save(&userChangeRight)

		ctx.JSON(http.StatusOK,gin.H{"status":"success"})
		return
	}else {
		ctx.JSON(http.StatusUnauthorized,gin.H{"error":"you are not an admin"})
		return
	}
}