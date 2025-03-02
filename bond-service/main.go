package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pratyush934/crowdfunding-app/bond-service/dbBond"
	"github.com/pratyush934/crowdfunding-app/bond-service/models"

	"net/http"
)

func loadDB() {
	dbBond.InitDB()

	dbBond.DB.AutoMigrate(&models.Bond{})
	dbBond.DB.AutoMigrate(&models.Transaction{})

}

func main() {

	router := gin.Default()

	//loadDB()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "It is working fine, I am Pratyush",
		})
	})

	router.Run("localhost:8081")
	fmt.Println("Server is running at port 8081")
}
