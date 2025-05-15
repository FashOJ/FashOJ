package router

import (
	"FashOJ_Backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func setupAuthRouter(r *gin.RouterGroup) {

	auth := r.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}
}
