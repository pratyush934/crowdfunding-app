package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pratyush934/crowdfunding-app/user-service/database"
	"github.com/pratyush934/crowdfunding-app/user-service/models"
)

func loadDB() {
	database.InitDB()
	err := database.DB.AutoMigrate(&models.User{})

	if err != nil {
		panic("Issue exist in Migrating User")
	}
}

func main() {
	fmt.Println("Hello I am Pratyush")

	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "I am Pratyush",
		})
	})

	loadDB()

	router.Run("localhost:8080")
}
