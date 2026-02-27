package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"idea-garden.tech/services/database"
	auth "idea-garden.tech/services/handlers/auth"
)


func main() {
	database.InitDB()
	response := gin.Default()
	response.Use(gin.Logger())
	response.Use(cors.Default())

	authRoutes := response.Group("/api/auth")
	{
		authRoutes.POST("/register", auth.Register)
		authRoutes.POST("/login", auth.Login)
	}

	response.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Fatal(response.Run(":8080"))
}