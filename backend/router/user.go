package router

import (
	"FashOJ_Backend/controllers"

	"github.com/gin-gonic/gin"
)

func setupUserRoutes(r *gin.RouterGroup) {
	
	user := r.Group("/user")
	{
		user.POST("/changepermission", controllers.ChangePermission)
	}
}