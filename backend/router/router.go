package router

import (
	"FashOJ_Backend/controllers"

	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine{
	r := gin.Default()

	auth := r.Group("/api/auth")
	{
		auth.POST("/login",controllers.Login)
		auth.POST("/register",controllers.Register)
	}

	return r
}