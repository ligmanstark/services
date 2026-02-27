package pkg

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/subosito/gotenv"
	er "idea-garden.tech/services/pkg/error"
)

type Claims struct {
	UserID int `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func (claims *Claims) CreateToken() (string, error) {
	err := gotenv.Load()
	er.HandleError("Ошибка в чтении .env", err)
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	expireTimeStr := os.Getenv("JWT_EXPIRE_TIME")

	if len(jwtSecret) == 0 {
		er.HandleError("JWT_SECRET не установлен в .env", err)
	}

	var expirationTime time.Time

	if expireTimeStr != ""{
		expireDuration,err := time.ParseDuration(expireTimeStr)
		if err != nil {
			er.HandleError("Ошибка в формате JWT_EXPIRE_TIME", err)
		}
		expirationTime = time.Now().Add(expireDuration)
	}

	claims.RegisteredClaims = jwt.RegisteredClaims{
		Subject: fmt.Sprintf("%d", claims.UserID),
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenStr, err := token.SignedString(jwtSecret)
	er.HandleError("Ошибка при создании JWT", err)

	return tokenStr, nil
}


func ValidateToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}

 	err := gotenv.Load()
	if err != nil {
 		fmt.Println("Предупреждение: ошибка в чтении .env:", err)
	}

 	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

 	if len(jwtSecret) == 0 {
		return nil, fmt.Errorf("JWT_SECRET не установлен в .env")
	}

 	token, err := jwt.ParseWithClaims(
		tokenStr,
		claims,
		func(t *jwt.Token) (interface{}, error) {
 
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf(
					"неожиданный алгоритм подписи: %v",
					t.Header["alg"],
				)
			}
			return jwtSecret, nil
		},
	)

 	if err != nil {
		return nil, fmt.Errorf("ошибка при валидации JWT: %w", err)
	}

 	if !token.Valid {
		return nil, fmt.Errorf("токен невалиден")
	}

 	return claims, nil
}
