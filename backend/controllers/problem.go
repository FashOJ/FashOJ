package controllers

import (
	"FashOJ_Backend/global"
	"FashOJ_Backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrUpdataProblem(ctx *gin.Context){
	var problem models.Problem
	if err := ctx.ShouldBindJSON(&problem); err != nil{
		ctx.JSON(http.StatusBadRequest,err.Error())
		return
	}

	if err := global.DB.AutoMigrate(&models.Problem{},&models.Example{},&models.Limit{}); err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	if err := global.DB.Save(&problem).Error;err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return		
	}

	ctx.JSON(http.StatusOK,gin.H{"status":"success"})
}