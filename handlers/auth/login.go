package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"idea-garden.tech/services/database"
	"idea-garden.tech/services/models"
	t "idea-garden.tech/services/pkg/token"
)

type LoginRequest struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(ctx *gin.Context){
var req LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Неверные данные для входа"})
		return
	}

	var user models.User
	db, err := database.InitDB()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Ошибка при подключении к базе данных"})
		return
	}
	defer db.Close()

	err = db.QueryRow("SELECT id, name, email, password FROM users WHERE email = $1", req.Email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		ctx.JSON(403, gin.H{"error": "Неверный email или пароль"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		ctx.JSON(403, gin.H{"error": "Неверный email или пароль"})
		return
	}

	claims := t.Claims{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
	}

	tokenStr, err := claims.CreateToken()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Ошибка при создании токена"})
		return
	}

	ctx.JSON(200, gin.H{"token": tokenStr})
}