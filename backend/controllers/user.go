package controllers

import (
	"FashOJ_Backend/global"
	"FashOJ_Backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChangePermission(ctx *gin.Context) {
	// Get the user from the context.
	user := ctx.Value("user").(models.User)

	// Get the username and OldPermission from the request body.
	var UserPatchPermission struct {
		Username   string
		Permission int
	}

	// Bind the request body to the UserPatchPermission struct.
	if err := ctx.ShouldBindJSON(&UserPatchPermission); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
		return
	}

	// Check if the user is an admin.
	if user.Permission == global.AdminUser {
		var userOldPermission models.User

		// Check if the user exists.
		if err := global.DB.Where("username = ?", UserPatchPermission.Username).First(&userOldPermission).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": "no such User or" + err.Error(),
			})
			return
		}

		// Check if the Permission code is valid.
		if UserPatchPermission.Permission < global.MinUserPermissionCode || UserPatchPermission.Permission > global.MaxUserPermissionCode {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": "wrong right code",
			})
			return
		}

		// Change the Permission code and save.
		userOldPermission.Permission = UserPatchPermission.Permission
		global.DB.Save(&userOldPermission)

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
