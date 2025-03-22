package middlewares

import (
	"FashOJ_Backend/global"
	"FashOJ_Backend/models"
	"FashOJ_Backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "no token"})
			ctx.Abort()
			return
		}

		username, err := utils.ParseJwt(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
			ctx.Abort()
			return
		}
		var user models.User
		if err := global.DB.Where("username = ?", username).First(&user).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
