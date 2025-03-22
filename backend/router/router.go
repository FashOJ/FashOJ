package router

import (
	"FashOJ_Backend/controllers"
	"FashOJ_Backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	user := r.Group("/api/user")
	user.Use(middlewares.AuthMiddleware())
	{
		user.POST("/changeright", controllers.ChangeRight)
	}

	normalProblem := r.Group("/api/problem")
	normalProblem.Use(middlewares.AuthMiddleware())
	{
		normalProblem.POST("",controllers.CreateOrUpdataProblem)
		normalProblem.POST("/:pid/upload",controllers.UploadTestcase)
	}

	normalSubmit := r.Group("/api/submit")
	normalSubmit.Use(middlewares.AuthMiddleware())

	return r
}
