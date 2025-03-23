package router

import (
	"FashOJ_Backend/controllers"
	"FashOJ_Backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	fashOJBackendRouter := gin.Default()

	auth := fashOJBackendRouter.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	user := fashOJBackendRouter.Group("/api/user")
	user.Use(middlewares.AuthMiddleware())
	{
		user.POST("/changepermission", controllers.ChangePermission)
	}

	normalProblem := fashOJBackendRouter.Group("/api/problem")
	normalProblem.Use(middlewares.AuthMiddleware())
	{
		normalProblem.POST("", controllers.CreateProblem)
		normalProblem.POST("/:pid/upload", controllers.UploadTestcase)
	}

	normalSubmit := fashOJBackendRouter.Group("/api/submit")
	normalSubmit.Use(middlewares.AuthMiddleware())

	return fashOJBackendRouter
}
