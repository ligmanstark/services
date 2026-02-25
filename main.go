package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"idea-garden.tech/services/database"
)


func main() {
	database.InitDB()
	response := gin.Default()
	response.Use(gin.Logger())
	response.Use(cors.Default())

	auth := response.Group("/api/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	response.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Fatal(response.Run(":8080"))
}