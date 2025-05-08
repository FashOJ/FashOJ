package controllers

import (
	"FashOJ_Backend/dto"
	"FashOJ_Backend/global"
	"FashOJ_Backend/models"
	"FashOJ_Backend/permission"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAnnouncement(ctx *gin.Context) {
	var announcement models.Announcement

	if !permission.HasPermission(ctx.Value("user").(models.User), permission.CreateAnnouncement) {
		ctx.JSON(http.StatusForbidden, gin.H{"message": "Insufficient permissions"})
		return
	}

	if err := ctx.ShouldBindJSON(&announcement); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "wrong format"})
		return
	}

	if err := global.DB.Create(&announcement).Error; err != nil {
		global.Logger.Errorf("%v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "something was wrong"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func GetLatestAnnouncement(ctx *gin.Context){
	var latestAnnouncement models.Announcement
	var dto dto.Announcement
	
	if err :=global.DB.Last(&latestAnnouncement).Error;err != nil {
		global.Logger.Errorf("%v",err.Error())
		ctx.JSON(http.StatusInternalServerError,gin.H{"message":"something was wrong"})
		return
	}

	dto.Title = latestAnnouncement.Title
	dto.Abstract = latestAnnouncement.Abstract

	ctx.JSON(http.StatusOK,gin.H{"message":"success","data":dto})
}