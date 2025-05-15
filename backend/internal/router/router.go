package router

import (
	"FashOJ_Backend/internal/global"
	"FashOJ_Backend/internal/middlewares"

	"github.com/gin-gonic/gin"

	"time"

	ginzap "github.com/gin-contrib/zap"
)

func SetupRouter() *gin.Engine {
	fashOJBackendRouter := gin.Default()

	// global middlewares
	fashOJBackendRouter.Use(ginzap.Ginzap(global.Logger, time.RFC3339, true))
	fashOJBackendRouter.Use(ginzap.RecoveryWithZap(global.Logger, true))
	fashOJBackendRouter.Use(middlewares.Cors())

	api := fashOJBackendRouter.Group("/api")
	{
		// 需要认证的路由
		setupNoAuthRoutes(api)
		setupAuthRouter(api)

		// 需要认证的路由组
		authorized := api.Group("")
		authorized.Use(middlewares.AuthMiddleware())
		{
			setupUserRoutes(authorized)
			setupNormalProblemRoutes(authorized)
			setupNormalSubmitRoutes(authorized)
			setupAnnouncementRoutes(authorized)
		}
	}

	return fashOJBackendRouter
}
