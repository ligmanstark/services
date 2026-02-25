package middleware

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	"idea-garden.tech/services/pkg"
)


func AuthMiddleware() gin.HandlerFunc {
	err := gotenv.Load()
	pkg.HandleError("Ошибка в чтении .env", err)
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	return func (ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Отсутствует токен авторизации"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	}
}



