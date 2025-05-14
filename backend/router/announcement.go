package router

import (
	"FashOJ_Backend/controllers"

	"github.com/gin-gonic/gin"
)

func setupAnnouncementRoutes(r *gin.RouterGroup) {
	
	announcement := r.Group("/announcement")
	{
		announcement.POST("", controllers.CreateAnnouncement)
		announcement.GET("", controllers.GetAnnouncement)
	}
}