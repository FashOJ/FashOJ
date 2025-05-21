package router

import (
	"FashOJ/Judger/internal/controllers"

	"github.com/gin-gonic/gin"
)

func setupProblemRouter(r *gin.RouterGroup) {

	r.POST("/problem",controllers.CreateProblem)

}