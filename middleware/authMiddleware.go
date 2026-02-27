package middleware

import (
	"github.com/gin-gonic/gin"
	t "idea-garden.tech/services/pkg/token"
)


func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")
		if tokenStr == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Отсутствует токен авторизации"})
			return
		}

		claims, err := t.ValidateToken(tokenStr)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Недействительный токен авторизации"})
			return
		}

		ctx.Set("userID", claims.UserID)
		ctx.Set("userName", claims.Name)
		ctx.Set("userEmail", claims.Email)

		ctx.Next()
	}
}



