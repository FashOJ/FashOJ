package controllers

import (
	"FashOJ_Backend/dto"
	"FashOJ_Backend/global"
	"FashOJ_Backend/models"
	"FashOJ_Backend/permission"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateAnnouncement(ctx *gin.Context) {
	var announcement models.Announcement
	var req dto.CreateAnnouncement
	var user = ctx.Value("user").(models.User)

	if !permission.HasPermission(user, permission.CreateAnnouncement) {
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

	if err := global.DB.Model(&user).Association("Announcement").Append(&announcement); err != nil {
		global.Logger.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "something was wrong"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func GetAnnouncement(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad query"})
		return
	}
	size, err := strconv.Atoi(ctx.DefaultQuery("size", "10"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad query"})
		return
	}

	var announcements []models.Announcement
	var total int64

	// 查询总数
	if err := global.DB.Model(&models.Announcement{}).Count(&total).Error; err != nil {
		global.Logger.Sugar().Errorf("%v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
		return
	}

	// 计算总页数
	totalPages := int(total) / size
	if int(total)%size != 0 {
		totalPages++
	}

	// 查询公告及关联的用户信息
	if err := global.DB.Preload("User").
		Offset((page - 1) * size).
		Limit(size).
		Order("created_at desc").
		Find(&announcements).Error; err != nil {
		global.Logger.Sugar().Errorf("%v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
		return
	}

	// 转换为DTO
	announcementDTOs := make([]dto.Announcement, 0, len(announcements))
	for _, a := range announcements {
		announcementDTOs = append(announcementDTOs, dto.Announcement{
			Author:   a.User.Username,
			Avatar:   a.User.Avatar,
			AuthorId: int(a.UserID),
			Id: int(a.ID),
			Title:    a.Title,
			Abstract: a.Abstract,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": &dto.AnnouncementPage{
			Pages:         totalPages,
			Size:          size,
			Announcements: announcementDTOs,
		},
	})
}

func GetLatestAnnouncement(ctx *gin.Context) {
	var latestAnnouncement models.Announcement
	var dto dto.LastAnnouncement

	if err := global.DB.Last(&latestAnnouncement).Error; err != nil {
		global.Logger.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "something was wrong"})
		return
	}

	dto.Title = latestAnnouncement.Title
	dto.Abstract = latestAnnouncement.Abstract

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": dto})
}

func abstractContent(content *string) string {
	var abstract string

	prefix, _, _ := strings.Cut(*content, "\n\n")
	abstract = prefix

	return abstract
}
