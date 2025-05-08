package controllers

import (
	"FashOJ_Backend/dto"
	"FashOJ_Backend/global"
	"FashOJ_Backend/models"
	"FashOJ_Backend/permission"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateAnnouncement(ctx *gin.Context) {
	var announcement models.Announcement
	var req dto.CreateAnnouncement

	if !permission.HasPermission(ctx.Value("user").(models.User), permission.CreateAnnouncement) {
		ctx.JSON(http.StatusForbidden, gin.H{"message": "Insufficient permissions"})
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "wrong format"})
		return
	}

	announcement.Title = req.Title
	announcement.Content = req.Content
	announcement.Abstract = abstractContent(&req.Content)

	if err := global.DB.Create(&announcement).Error; err != nil {
		global.Logger.Errorf("%v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "something was wrong"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func GetLatestAnnouncement(ctx *gin.Context){
	var latestAnnouncement models.Announcement
	var dto dto.LastAnnouncement
	
	if err :=global.DB.Last(&latestAnnouncement).Error;err != nil {
		global.Logger.Errorf("%v",err.Error())
		ctx.JSON(http.StatusInternalServerError,gin.H{"message":"something was wrong"})
		return
	}

	dto.Title = latestAnnouncement.Title
	dto.Abstract = latestAnnouncement.Abstract

	ctx.JSON(http.StatusOK,gin.H{"message":"success","data":dto})
}


func abstractContent(content *string) string{
	var abstract string

	prefix,_,_:=strings.Cut(*content,"\n\n");
	abstract = prefix

	return abstract
}