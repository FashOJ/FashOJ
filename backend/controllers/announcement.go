package controllers

import (
	"FashOJ_Backend/global"
	"FashOJ_Backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAnnouncement(ctx *gin.Context) {
	var announcement models.Announcement

	if err := ctx.ShouldBindJSON(&announcement); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "wrong format"})
		return
	}

	if err := global.DB.Create(&announcement).Error; err != nil {
		global.Logger.Errorf("%v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "something was wrong"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Status": "succes"})
}
