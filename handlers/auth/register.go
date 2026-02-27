package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"idea-garden.tech/services/database"
	"idea-garden.tech/services/models"
	t "idea-garden.tech/services/pkg/token"
)

type RegisterRequest struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func Register(ctx *gin.Context){
	var req RegisterRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Неверные данные для регистрации"})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Ошибка при обработке пароля"})
		return
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashPassword),
	}

	db,err := database.InitDB()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Ошибка при подключении к базе данных"})
		return
	}
	defer db.Close()

	err = db.QueryRow("INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id",
		user.Name, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Ошибка при сохранении пользователя"})
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

	ctx.JSON(201, gin.H{"token": tokenStr})
 }