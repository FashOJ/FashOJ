package router

import (
	"FashOJ_Backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func setupNormalProblemRoutes(r *gin.RouterGroup) {

	normalProblem := r.Group("/problem")
	{
		normalProblem.POST("", controllers.CreateProblem)
		normalProblem.POST("/:pid/upload", controllers.UploadTestcase)
	}
}
