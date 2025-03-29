package controllers

import (
	"FashOJ_Backend/global"
	"FashOJ_Backend/models"
	"FashOJ_Backend/permission"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChangePermission(ctx *gin.Context) {
	// Get the user from the context.
	user := ctx.Value("user").(models.User)

	// Get the username and OldPermission from the request body.
	var UserPatchPermission struct {
		Username   string `binding:"required"`
		Permission uint32 `binding:"required"`
	}

	// Bind the request body to the UserPatchPermission struct.
	if err := ctx.ShouldBindJSON(&UserPatchPermission); err != nil {
		global.Logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
		return
	}

	// Check if the user is has permission to change others permission.
	if permission.HasPermission(user, permission.ModifyPermission) {
		var userOldPermission models.User

		// Check if the user exists.
		if err := global.DB.Where("username = ?", UserPatchPermission.Username).First(&userOldPermission).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": "no such User or" + err.Error(),
			})
			return
		}

		// Check if the Permission code is valid.
		if permission.IsVaild(UserPatchPermission.Permission) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": "wrong right code",
			})
			return
		}

		// Change the Permission code and save.
		userOldPermission.Permission = UserPatchPermission.Permission
		if err := global.DB.Save(&userOldPermission).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Error": "wrong",
			})
			global.Logger.Errorln(err.Error())
			return
		}

		// Return success.
		ctx.JSON(http.StatusOK, gin.H{
			"Status": "success",
		})
		return
	} else {
		// If the user is not an admin, return unauthorized.
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Error": "You are not admin, and have no permission to patch.",
		})
		return
	}
}
