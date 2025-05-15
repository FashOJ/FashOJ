package router

import (
	"FashOJ_Backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func setupNoAuthRoutes(r *gin.RouterGroup) {
	noAuth := r.Group("")
	{
		noAuth.GET("/announcement/latest", controllers.GetLatestAnnouncement)
	}

}
