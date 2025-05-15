package middlewares

import (
	"FashOJ_Backend/internal/global"
	"FashOJ_Backend/internal/models"
	"FashOJ_Backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleware is a middleware that checks if the user is authenticated.
// If the user is authenticated, the middleware will set the user in the context.
// If the user is not authenticated, the middleware will return an error.
func AuthMiddleware() gin.HandlerFunc {

	// Check if the token is valid
	return func(ctx *gin.Context) {

		// Get the token from the header
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": "No token",
			})
			ctx.Abort()
			return
		}

		// Check if the token is valid.
		username, err := utils.ParseJwt(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
			ctx.Abort()
			return
		}

		// Check if the user exists.
		var authenticatingUser models.User
		if err := global.DB.Where("username = ?", username).First(&authenticatingUser).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Error": err,
			})
			ctx.Abort()
			return
		}
		ctx.Set("user", authenticatingUser)
		ctx.Next() // Continue to the next middleware
	}
}
