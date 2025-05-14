package router

import (
	"FashOJ_Backend/controllers"

	"github.com/gin-gonic/gin"
)

func setupNoAuthRoutes(r *gin.RouterGroup) {
	noAuth := r.Group("")
	{
		noAuth.GET("/announcement/latest", controllers.GetLatestAnnouncement)
	}

}
