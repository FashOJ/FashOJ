package controllers

import (
	"FashOJ_Backend/global"
	"FashOJ_Backend/models"
	"FashOJ_Backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context){

	var input models.User
	if err:=ctx.ShouldBindJSON(&input);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"error":err,})
		return
	}

	var user models.User
	if err := global.DB.Where("username = ?",input.Username).First(&user).Error;err != nil{
		ctx.JSON(http.StatusUnauthorized,gin.H{"error":"wrong password or username"})
		return
	}

	if !utils.CheckPwd(input.Password,user.Password){
		ctx.JSON(http.StatusUnauthorized,gin.H{"error":"wrong password or username"})
		return
	}

	token,err := utils.GenJwt(user.Username)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{"error":err})
	}

	ctx.JSON(http.StatusOK,gin.H{"token":token})

}

func Register(ctx *gin.Context){
	var user models.User
	if err:=ctx.ShouldBindJSON(&user);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"error":err,})
		return
	}

	hashedPwd,err := utils.HashPwd(user.Password)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{"error":err,})
	}

	user.Password = hashedPwd
	token,err := utils.GenJwt(user.Username)
	if err !=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{"error":err,})
		return
	}

	if err:=global.DB.AutoMigrate(&user); err!=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{"error":err,})
		return
	}
	
	if err := global.DB.Create(&user).Error;err!=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{"error":err})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{"token":token})
}