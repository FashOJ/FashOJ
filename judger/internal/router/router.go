package router

import "github.com/gin-gonic/gin"

func SetupRouter(r *gin.RouterGroup) {

	judger := r.Group("/judger")

	setupProblemRouter(judger)
}