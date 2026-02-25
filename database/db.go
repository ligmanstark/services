package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/subosito/gotenv"
	"idea-garden.tech/services/pkg"
)

func InitDB()(*sql.DB, error){
	err:=gotenv.Load()
	pkg.HandleError("Ошибка в чтении .env",err)

	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
		sslmode  = os.Getenv("DB_SSLMODE")
	)
	portInt, err := strconv.Atoi(port)
	pkg.HandleError("Ошибка в конвертации порта из строки в число", err)

	return sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, portInt, user, password, dbname, sslmode))
}