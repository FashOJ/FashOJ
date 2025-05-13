package router

import (
	"FashOJ_Backend/controllers"
	"FashOJ_Backend/global"
	"FashOJ_Backend/middlewares"

	"github.com/gin-gonic/gin"

	"time"

	ginzap "github.com/gin-contrib/zap"
)

func SetupRouter() *gin.Engine {
	fashOJBackendRouter := gin.Default()

	fashOJBackendRouter.Use(ginzap.Ginzap(global.Logger, time.RFC3339, true))
	fashOJBackendRouter.Use(ginzap.RecoveryWithZap(global.Logger, true))

	fashOJBackendRouter.Use(middlewares.Cors())

	noAuth := fashOJBackendRouter.Group("")
	{
		noAuth.GET("/api/announcement/latest", controllers.GetLatestAnnouncement)
	}

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

	announcement := fashOJBackendRouter.Group("/api/announcement")
	announcement.Use(middlewares.AuthMiddleware())
	{
		announcement.POST("", controllers.CreateAnnouncement)
		announcement.GET("",controllers.GetAnnouncement)
	}
	return fashOJBackendRouter
}
