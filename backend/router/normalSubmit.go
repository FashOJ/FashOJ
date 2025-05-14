package router

import "github.com/gin-gonic/gin"

func setupNormalSubmitRoutes(r *gin.RouterGroup) {
	
	normalSubmit := r.Group("/submit")
	normalSubmit.POST("")
}