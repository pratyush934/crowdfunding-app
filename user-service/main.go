package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pratyush934/crowdfunding-app/user-service/controller"
	"github.com/pratyush934/crowdfunding-app/user-service/database"
	"github.com/pratyush934/crowdfunding-app/user-service/models"
	"github.com/pratyush934/crowdfunding-app/user-service/util"
)

func loadDB() {
	database.InitDB()
	err := database.DB.AutoMigrate(&models.User{})
	err1 := database.DB.AutoMigrate(&models.Role{})

	if err != nil {
		fmt.Println("There exist an error while migrating User")
		return
	}

	if err1 != nil {
		fmt.Println("There exist an error while migrating Role")
		return
	}
	seedData()
}

func seedData() {
	var roles = []models.Role{
		{Name: "admin", Description: "You are head for all the things", ID: 1},
		{Name: "user", Description: "You are the user", ID: 2},
	}
	var users = []models.User{
		{UserName: "Pratyush", Email: "pratyush@gmail.com", Password: "12345", RoleId: 1, BondID: "[]"},
		{UserName: "Rahul", Email: "rahulwa@gmail.com", Password: "67890", RoleId: 2, BondID: "[]"},
	}

	if err := database.DB.Save(&roles).Error; err != nil {
		fmt.Println("Error saving roles:", err)
	}
	for _, user := range users {
		if err := database.DB.Save(&user).Error; err != nil {
			fmt.Println("Error saving user:", user.UserName, err)
		}
	}
}

func main() {
	fmt.Println("Hello I am Pratyush")

	router := gin.Default()
	loadDB()

	authRoutes := router.Group("/auth")
	authRoutes.POST("/register", controller.Register)
	authRoutes.POST("/login", controller.Login)

	adminRoutes := router.Group("/admin")
	adminRoutes.Use(util.JWTAuth())
	adminRoutes.GET("/users", controller.GetUsers)
	adminRoutes.GET("/users/:id", controller.GetUser)
	adminRoutes.PUT("/users/:id", controller.UpdateUser)
	adminRoutes.POST("/users/role", controller.AssignRole)
	adminRoutes.GET("/users/roles", controller.GetRoles)
	adminRoutes.GET("/users/roles/:id", controller.GetRole)
	adminRoutes.PUT("/users/roles/:id", controller.UpdateRole)

	router.Run("localhost:8080")
	fmt.Println("Server running at the port 8080")
}
